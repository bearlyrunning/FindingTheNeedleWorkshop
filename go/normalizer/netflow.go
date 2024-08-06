package main

import (
	"log"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
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

	// <TODO: Implement me!>
	// Implement the missing validate function(s) in validator.go file.
	// Parse and return `datetime` field with validateTime().
	timestamp, err := validateTime("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `src_ip` field with validateIP().
	src_ip, err := validateIP("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `src_port` field with validatePort().
	src_port, err := validatePort("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `dst_ip` field with validateIP().
	dst_ip, err := validateIP("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `dst_port` field with validatePort().
	dst_port, err := validatePort("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `bytes_in` field with validateInt64().
	bytes_in, err := validateInt64("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `bytes_out` field with validateInt64().
	bytes_out, err := validateInt64("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `packets_in` field with validateInt64().
	packets_in, err := validateInt64("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `packets_out` field with validateInt64().
	packets_out, err := validateInt64("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// Return a populated NormalizedLog proto message.
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
