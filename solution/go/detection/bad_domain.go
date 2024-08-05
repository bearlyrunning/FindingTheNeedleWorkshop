package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

const indicatorPath = "../../data/indicators/bad_domain.csv"

func (bdd *BadDomainDetection) ruleName() string {
	return bdd.name
}

func fmtRegex(ind []string) string {
	return fmt.Sprintf(".*(%s)$", strings.Join(ind, "|"))
}

func (bdd *BadDomainDetection) setFilterRegex() error {
	// Get the list of domain indicators.
	var ind []string
	f, err := os.Open(indicatorPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", indicatorPath, err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		ind = append(ind, strings.Split(s.Text(), ",")[0])
	}
	if err = s.Err(); err != nil {
		return fmt.Errorf("failed to scan file: %v", err)
	}
	// Compile a regex expression for matching indicators of compromise

	str := fmtRegex(ind)
	bdd.rr, err = regexp.Compile(str)
	if err != nil {
		return fmt.Errorf("failed compiling regex %s: %v", str, err)
	}
	return nil
}

// Filter out DNS logs containing known bad domains.
func (bdd *BadDomainDetection) filter() []*nlpb.DNS {
	var matched []*nlpb.DNS
	for _, log := range bdd.logs.dns {
		if bdd.rr.MatchString(log.GetQuery()) || bdd.rr.MatchString(log.GetAnswer()) {
			matched = append(matched, log)
		}
	}
	return matched
}

// Aggregate DNS logs containing known bad domains by source IP
// along with the time window of observed bad traffic.
func (bdd *BadDomainDetection) aggregate(matched []*nlpb.DNS) []*spb.Signal {
	var sigs []*spb.Signal
	keyBySrc := make(map[string][]*nlpb.DNS)
	for _, m := range matched {
		keyBySrc[m.SourceIp] = append(keyBySrc[m.SourceIp], m)
	}
	for src, logs := range keyBySrc {
		// Identify the time window of the bad DNS traffic.
		earliest, latest := logs[0].Timestamp.AsTime(), logs[0].Timestamp.AsTime()
		for _, l := range logs[1:] {
			if l.Timestamp.AsTime().Before(earliest) {
				earliest = l.Timestamp.AsTime()
				continue
			}
			if l.Timestamp.AsTime().After(latest) {
				latest = l.Timestamp.AsTime()
			}
		}

		// Find the bad indicator queried.
		m := bdd.rr.FindStringSubmatch(logs[0].GetQuery())
		if len(m) == 0 {
			m = bdd.rr.FindStringSubmatch(logs[0].GetAnswer())
		}

		sigs = append(sigs, &spb.Signal{
			Event: &spb.Signal_BadDomain{
				BadDomain: &spb.BadDomain{
					TimestampStart: tspb.New(earliest),
					TimestampEnd:   tspb.New(latest),
					SourceIp:       src,
					BadDomain:      m[1],
					DnsLog:         logs,
				},
			},
		})
	}
	return sigs
}

func (bdd *BadDomainDetection) run() ([]*spb.Signal, error) {
	// Set regex for filtering.
	if err := bdd.setFilterRegex(); err != nil {
		return nil, err
	}
	// Run detection logic: 1) filter 2) aggregate.
	return bdd.aggregate(bdd.filter()), nil
}
