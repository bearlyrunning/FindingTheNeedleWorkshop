package main

import (
	"log"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
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
	timestamp, err := validateTime(fields[0])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	src_ip, err := validateIP(fields[2])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	resolver_ip, err := validateIP(fields[3])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	query, err := validateQuery(fields[4])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	code, err := validateReturnCode(fields[7])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}

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
