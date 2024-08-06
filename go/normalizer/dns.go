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

// normalize() function does the following:
// - reads each `line` of /data/dns/dns.log
// - parses and validates each comma separated field in the log
// - output the log as a NormalizedLog proto message which eventually get saved as:
//   - /data/dns/dns.binpb
//   - /data/dns/dns.json
func (dn *DNSNormalizer) normalize(line string) *nlpb.NormalizedLog {
	fields := strings.Split(line, ",")

	// Validate number of fields in each log line.
	if len(fields) != 8 {
		log.Printf("invalid number of fields found; expect 8, found %d: %s\n", len(fields), line)
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
	// Parse and return `resolver_ip` field with validateIP().
	resolver_ip, err := validateIP("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `query` field with validateQuery().
	query, err := validateQuery("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// <TODO: Implement me!>
	// Parse and return `return_code` field with validateReturnCode().
	code, err := validateReturnCode("TODO: Implement me!")
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

	// Return a populated NormalizedLog proto message.
	return &nlpb.NormalizedLog{
		Msg: &nlpb.NormalizedLog_DnsLog{
			DnsLog: &nlpb.DNS{
				Timestamp:  timestamp,
				SourceIp:   src_ip,
				ResolverIp: resolver_ip,
				Query:      query,
				Type:       fields[5],
				Answer:     fields[6],
				ReturnCode: code,
				LogSource:  fields[1],
			},
		},
	}
}
