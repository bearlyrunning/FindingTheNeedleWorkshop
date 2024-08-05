package main

import (
	"fmt"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

const timeFmt = "2006-01-02 15:04:05.000"

func validateTime(s string) (*tspb.Timestamp, error) {
	// <TODO: Implement me!>
	// Parse a datetime string to a Timestamp proto message.
	// Return error fmt.Errorf("invalid timestamp found: %s", s) if the datetime string is not valid.
	// Hint #1: import "time".
	// Hint #2: make use of the timestamppb package.
	return nil, nil
}

func validateTimestamp(s string) (*tspb.Timestamp, error) {
	// <TODO: Implement me!>
	// Parse a epoch timestamp string to a Timestamp proto message.
	// Return error fmt.Errorf("unexpected timestamp found: %s", s) if the timestamp string is not valid.
	// Fix the placeholder return below.
	// Hint #1: what field(s) does the Timestamp proto message contain?
	return nil, nil
}

func validateIP(s string) (string, error) {
	// <TODO: Implement me!>
	// Confirm the IP string contains a valid IP address.
	// Return back the valid IP.
	// Return error fmt.Errorf("invalid IP found: %s", s) if the IP is not valid.
	return "", nil
}

func validatePort(s string) (int32, error) {
	// <TODO: Implement me!>
	// Convert port strings to int32.
	// Confirm the string is a valid port number.
	// Return error fmt.Errorf("unexpected port number found: %s", s) if the port is not valid.
	return int32(0), nil
}

func validateQuery(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("empty query found")
	}
	return s, nil
}

func validateReturnCode(s string) (nlpb.DNS_ReturnCode, error) {
	// <TODO: Implement me!>
	// Convert return code to proto ENUM nlpb.DNS_ReturnCode.
	// Confirm the string is a valid return code (hint: check the range of the enum).
	// Return error fmt.Errorf("unexpected return code found: %s", s) if the code is not valid.
	// Don't forget to increment return code by 1 as the enum value 0 is reserved for default value (e.g. unspecified) only.
	// Hint: check the auto-generated normalizedlogpb package for suitable conversion approach.
	return nlpb.DNS_UNSPECIFIED, nil
}

func validateInt64(s string) (int64, error) {
	// <TODO: Implement me!>
	// Convert number string to int64.
	// Return error fmt.Errorf("unexpected string found, expecting int64: %s", s) if the string is not a valid int64.
	return 0, nil
}

func validatePlatform(s string) nlpb.Execution_Platform {
	// <TODO: Implement me!>
	// Convert platform string to Platform ENUM.
	// If the platform string is not valid, set to default nlpb.Execution_Platform(0).
	// Hint: check the auto-generated normalizedlogpb package for suitable conversion approach.
	return nlpb.Execution_UNSPECIFIED
}
