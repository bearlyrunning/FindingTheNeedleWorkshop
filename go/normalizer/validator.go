package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

const timeFmt = "2006-01-02 15:04:05.000"

func validateTime(s string) (*tspb.Timestamp, error) {
	t, err := time.Parse(timeFmt, s)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp found: %s", s)
	}
	// Other checks that could be implemented: check if timestamp is not in the future or too far in the past.
	return tspb.New(t), nil
}

func validateTimestamp(s string) (*tspb.Timestamp, error) {
	ts, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("unexpected timestamp found: %s", s)
	}
	// Other checks that could be implemented: check if timestamp is not in the future or too far in the past.
	return &tspb.Timestamp{Seconds: ts}, nil
}

func validateIP(s string) (string, error) {
	// <TODO: Implement me!>
	// Confirm the IP string contains a valid IP address.
	// Return back the valid IP.
	// Return error fmt.Errorf("invalid IP found: %s", s) if the IP is not valid.
	return "", nil
}

func validatePort(s string) (int32, error) {
	var port int32
	// <TODO: Implement me!>
	// Convert port strings to int32.
	// Confirm the string is a valid port number.
	// Return error fmt.Errorf("unexpected port number found: %s", s) if the port is not valid.
	return port, nil
}

func validateQuery(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("empty query found")
	}
	return s, nil
}

func validateReturnCode(s string) (nlpb.DNS_ReturnCode, error) {
	code, err := strconv.ParseInt(s, 10, 32)
	if err != nil || code > 9 {
		return 0, fmt.Errorf("unexpected return code found: %s", s)
	}
	// Increment return code by 1 as the enum value 0 is reserved for default value (e.g. unspecified) only.
	return nlpb.DNS_ReturnCode(int32(code + 1)), nil
}

func validateInt64(s string) (int64, error) {
	var i int64
	// <TODO: Implement me!>
	// Convert number string to int64.
	// Return error fmt.Errorf("unexpected string found, expecting int64: %s", s) if the string is not a valid int64.
	return i, nil
}

func validatePlatform(s string) nlpb.Execution_Platform {
	p, ok := nlpb.Execution_Platform_value[strings.Trim(s, "\"")]
	if !ok {
		log.Printf("invalid platform %s found, set to default", s)
		return nlpb.Execution_Platform(0)
	}
	return nlpb.Execution_Platform(p)
}
