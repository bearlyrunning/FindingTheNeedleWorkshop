package main

import (
	"fmt"

	spb "github.com/bearlyrunning/FindingTheNeedle/go/generated/signalpb"
)

// Threshold of total number of bytes for identifying DNS tunnelling.
const threshold = 480

func (dtd *DNSTunnelDetection) ruleName() string {
	return dtd.name
}

// run() function does the following:
//   - loop through each proto in the /data/netflow/netflow_normalized.binpb or json file
//     (these protos are saved in bdd.logs.netflow)
//   - apply detection logic
//   - output the log as a Signal proto (see /proto/signal.proto) message which
//     eventually get saved in /data/signal/dns_tunnel.json
func (dtd *DNSTunnelDetection) run() ([]*spb.Signal, error) {
	// <TODO: Implement me!>
	// Find relevant netflow logs indicating potential DNS tunneling behaviour:
	// To simplify the logic, the rule contains the following steps:
	//   1. Filter logs to what is relevant, then
	//   2. Aggregate logs based on source IP-destination IP address pairs.
	//   3. Only return logs with aggregated outbound traffic volume above THRESHOLD.
	//   4. Return the set of interesting logs as a list of spb.Signal.

	// Expected output: the list of spb.Signal returned should have `event` field set to `dns_tunnel`.

	// Hint #1: Assume DNS traffic is on port 53 inbound and outbound.
	// Hint #2: In your map, construct a key using the values of source and destination IP addresses (for your pairing).
	// Hint #3: We need to compare traffic volume against a `threshold` - remember to keep a sum of bytes_in and bytes_out to check if we exceed the threshold.
	// Hint #3: Check the fields you need to populate by inspecting the spb.DNSTunnel protobuf message.

	var sigs []*spb.Signal

	for _, n := range dtd.logs.netflow {
		fmt.Printf("TODO: Implement me! %v", n)
	}

	sigs = append(sigs, &spb.Signal{
		Event: &spb.Signal_DnsTunnel{
			DnsTunnel: &spb.DNSTunnel{
				// TimestampStart: ,
				// TimestampEnd:   ,
				// SourceIp:       ,
				// TunnelIp:       ,
				// BytesInTotal:   ,
				// BytesOutTotal:  ,
				// NetflowLog:     logs,
			},
		},
	})

	return sigs, nil
}
