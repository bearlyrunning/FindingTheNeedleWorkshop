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
	// Implement the validate function in validator.go file.
	// Parse and return `datetime` field with validateTime().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `src_ip` field with validateIP().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `src_port` field with validatePort().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `dst_ip` field with validateIP().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `dst_port` field with validatePort().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `bytes_in` field with validateInt64().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `bytes_out` field with validateInt64().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `packets_in` field with validateInt64().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `packets_out` field with validateInt64().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Return a populated NormalizedLog proto message.
	return &nlpb.NormalizedLog{
		Msg: &nlpb.NormalizedLog_NetflowLog{},
	}
}
