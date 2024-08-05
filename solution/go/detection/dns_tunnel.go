package main

import (
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

// Threshold of total number of bytes for identifying DNS tunnelling.
const threshold = 480

func (dtd *DNSTunnelDetection) ruleName() string {
	return dtd.name
}

// Filter out netflow logs for DNS traffic.
func (dtd *DNSTunnelDetection) filter() []*nlpb.Netflow {
	var matched []*nlpb.Netflow
	for _, n := range dtd.logs.netflow {
		if n.GetDstPort() == 53 || n.GetSrcPort() == 53 {
			matched = append(matched, n)
		}
	}
	return matched
}

// Aggregate traffic by IP pairs and count number of bytes transported over DNS.
func (dtd *DNSTunnelDetection) aggregate(matched []*nlpb.Netflow) []*spb.Signal {
	var sigs []*spb.Signal

	// Parse netflow by grouping logs by IP pairs.
	keyByIpPair := make(map[string][]*nlpb.Netflow)
	for _, m := range matched {
		// Set key format: client IP, remote IP
		k := m.GetSrcIp() + "," + m.GetDstIp()
		// If getting response from remote IP, reverse the order.
		if m.GetSrcPort() == 53 {
			k = m.GetDstIp() + "," + m.GetSrcIp()
		}
		keyByIpPair[k] = append(keyByIpPair[k], m)
	}

	for pair, logs := range keyByIpPair {
		earliest, latest := logs[0].GetTimestamp().AsTime(), logs[0].GetTimestamp().AsTime()
		// Assuming bytes in and out numbers are set with directions of the traffic taken into consideration.
		in, out := logs[0].GetBytesIn(), logs[0].GetBytesOut()
		for _, l := range logs[1:] {
			in += l.GetBytesIn()
			out += l.GetBytesOut()
			if l.Timestamp.AsTime().Before(earliest) {
				earliest = l.Timestamp.AsTime()
				continue
			}
			if l.Timestamp.AsTime().After(latest) {
				latest = l.Timestamp.AsTime()
			}
		}
		// Identify potential DNS tunnelling by checking total number of bytes against a static threshold.
		if out > threshold {
			sigs = append(sigs, &spb.Signal{
				Event: &spb.Signal_DnsTunnel{
					DnsTunnel: &spb.DNSTunnel{
						TimestampStart: tspb.New(earliest),
						TimestampEnd:   tspb.New(latest),
						SourceIp:       strings.Split(pair, ",")[0],
						TunnelIp:       strings.Split(pair, ",")[1],
						BytesInTotal:   in,
						BytesOutTotal:  out,
						NetflowLog:     logs,
					},
				},
			})
		}
	}

	return sigs
}

func (dtd *DNSTunnelDetection) run() ([]*spb.Signal, error) {
	return dtd.aggregate(dtd.filter()), nil
}
