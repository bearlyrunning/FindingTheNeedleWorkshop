package main

import (
	"log"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
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

func (en *ExecutionNormalizer) normalize(line string) *nlpb.NormalizedLog {
	fields := splitWithEscape(line, ",", "\\")

	// Validate fields.
	if len(fields) != 9 {
		log.Printf("invalid number of fields found; expect 9, found %d: %s\n", len(fields), line)
		return nil
	}
	timestamp, err := validateTimestamp(fields[0])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	uid, err := validateInt64(fields[3])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	pid, err := validateInt64(fields[4])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	ppid, err := validateInt64(fields[5])
	if err != nil {
		log.Printf("%v, skipping: %s\n", err, line)
		return nil
	}
	platform := validatePlatform(fields[8])
	return &nlpb.NormalizedLog{
		Msg: &nlpb.NormalizedLog_ExecutionLog{
			ExecutionLog: &nlpb.Execution{
				Timestamp: timestamp,
				Filepath:  strings.Trim(fields[1], "\""),
				Command:   strings.Trim(fields[2], "\""),
				Uid:       uid,
				Pid:       pid,
				Ppid:      ppid,
				Cwd:       strings.Trim(fields[6], "\""),
				Hostname:  strings.Trim(fields[7], "\""),
				Platform:  platform,
			},
		},
	}
}
