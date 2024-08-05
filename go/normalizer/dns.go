package main

import (
	"log"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
)

func (dn *DNSNormalizer) getInput() string {
	return dn.input
}

func (dn *DNSNormalizer) getBinaryOutput() string {
	return dn.binaryOutput
}

func (dn *DNSNormalizer) getJsonOutput() string {
	return dn.jsonOutput
}

func (dn *DNSNormalizer) normalize(line string) *nlpb.NormalizedLog {
	fields := strings.Split(line, ",")

	// Validate fields.
	if len(fields) != 8 {
		log.Printf("invalid number of fields found; expect 8, found %d: %s\n", len(fields), line)
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
	// Parse and return `resolver_ip` field with validateIP().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `query` field with validateQuery().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `return_code` field with validateReturnCode().
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Return a populated NormalizedLog proto message.
	return &nlpb.NormalizedLog{}
}
