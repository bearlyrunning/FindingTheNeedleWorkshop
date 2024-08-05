package main

import (
	"log"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
)

func (nn *NetflowNormalizer) getInput() string {
	return nn.input
}

func (nn *NetflowNormalizer) getBinaryOutput() string {
	return nn.binaryOutput
}

func (nn *NetflowNormalizer) getJsonOutput() string {
	return nn.jsonOutput
}

func (nn *NetflowNormalizer) normalize(line string) *nlpb.NormalizedLog {
	fields := strings.Split(line, ",")

	// Validate fields.
	if len(fields) != 11 {
		log.Printf("invalid number of fields found; expect 11, found %d: %s\n", len(fields), line)
		return nil
	}
	timestamp, err := validateTime(fields[0])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	src_ip, err := validateIP(fields[3])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	src_port, err := validatePort(fields[4])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	dst_ip, err := validateIP(fields[5])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	dst_port, err := validatePort(fields[6])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	bytes_in, err := validateInt64(fields[7])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	bytes_out, err := validateInt64(fields[8])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	packets_in, err := validateInt64(fields[9])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	packets_out, err := validateInt64(fields[10])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	return &nlpb.NormalizedLog{
		Msg: &nlpb.NormalizedLog_NetflowLog{
			NetflowLog: &nlpb.Netflow{
				Timestamp:  timestamp,
				SrcIp:      src_ip,
				SrcPort:    src_port,
				DstIp:      dst_ip,
				DstPort:    dst_port,
				BytesIn:    bytes_in,
				BytesOut:   bytes_out,
				PacketsIn:  packets_in,
				PacketsOut: packets_out,
				Protocol:   fields[2],
				LogSource:  fields[1],
			},
		},
	}
}
