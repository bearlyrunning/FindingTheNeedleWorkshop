package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	spb "github.com/bearlyrunning/FindingTheNeedle/go/generated/signalpb"
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

// run() function does the following:
//   - loop through each proto in the /data/dns/dns_normalized.binpb or json file
//     (these protos are saved in bdd.logs.dns)
//   - apply detection logic
//   - output the log as a Signal proto (see /proto/signal.proto) message which
//     eventually get saved in /data/signal/bad_domain.json
func (bdd *BadDomainDetection) run() ([]*spb.Signal, error) {
	// Set regex for filtering.
	if err := bdd.setFilterRegex(); err != nil {
		return nil, err
	}

	// <TODO: Implement me!>
	// Find any logs that contain indicators of compromise from indicatorPath:
	//   1. Filter logs to what is relevant, then
	//   2. [Optional] Aggregate logs based on source IP address.
	//   3. Return the set of interesting logs as a list of spb.Signal

	// Expected output:
	// Option #1: If the aggregation step is skipped, the list of spb.Signal returned should have `event` field set to `bad_domain_filtered`.
	// Option #2: If both filtering and aggregation are performed, the list of spb.Signal returned should have `event` field set to `bad_domain`.

	// Hint #1: Make use of bdd.rr and the `regexp` package.
	// Hint #2: Aggregation is easier using a map data structure.
	// Hint #3: Check the fields you need to populate by inspecting the spb.BadDomain protobuf message.
	var sigs []*spb.Signal
	for _, log := range bdd.logs.dns {
		fmt.Printf("TODO: Implement me! %v", log)
	}

	sigs = append(sigs, &spb.Signal{
		Event: &spb.Signal_BadDomain{
			BadDomain: &spb.BadDomain{
				// TimestampStart: ,
				// TimestampEnd:   ,
				// SourceIp:       ,
				// BadDomain:      ,
				// DnsLog:         ,
			},
		},
	})
	return sigs, nil
}
