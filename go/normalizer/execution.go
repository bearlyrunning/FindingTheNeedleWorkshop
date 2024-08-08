package main

import (
	"log"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
)

func (en *ExecutionNormalizer) getInput() string {
	return en.input
}

func (en *ExecutionNormalizer) getBinaryOutput() string {
	return en.binaryOutput
}

func (en *ExecutionNormalizer) getJsonOutput() string {
	return en.jsonOutput
}

func splitWithEscape(str, sep, esc string) []string {
	strs := strings.Split(str, sep)

	var escapedStrs []string
	concat := false
	concatStr := ""
	for _, s := range strs {
		// If current element ends with "\", enable concatnation.
		if strings.HasSuffix(s, esc) {
			if !concat {
				concat = true
			}
			concatStr += strings.ReplaceAll(s, "\\", ",")
		} else {
			concatStr += s
			escapedStrs = append(escapedStrs, concatStr)
			concat = false
			concatStr = ""
		}
	}
	return escapedStrs
}

// normalize() function does the following:
// - reads each `line` of /data/execution/execution.log
// - parses and validates each comma separated field in the log
// - output the log as a NormalizedLog proto (see /proto/normalized.proto) message which eventually get saved as:
//   - /data/execution/execution_normalized.binpb
//   - /data/execution/execution_normalized.json
func (en *ExecutionNormalizer) normalize(line string) *nlpb.NormalizedLog {
	fields := splitWithEscape(line, ",", "\\")

	// Validate fields.
	if len(fields) != 9 {
		log.Printf("invalid number of fields found; expect 9, found %d: %s\n", len(fields), line)
		return nil
	}

	// <TODO: Implement me!> Implement the missing logic below.

	// <TODO: Implement me!>
	// Parse and return `timestamp` field with validateTimestamp().

	// <TODO: Implement me!>
	// Parse and return `uid` field with validateInt64().

	// <TODO: Implement me!>
	// Parse and return `pid` field with validateInt64().

	// <TODO: Implement me!>
	// Parse and return `ppid` field with validateInt64().

	// <TODO: Implement me!>
	// Parse and return `platform` field with validatePlatform().

	// <TODO: Implement me!>
	// Return a populated NormalizedLog proto message.
	return &nlpb.NormalizedLog{}
}
