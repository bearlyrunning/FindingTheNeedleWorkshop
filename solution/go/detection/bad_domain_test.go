package main

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
)

func TestFmtRegex(t *testing.T) {
	tc := struct {
		iocs []string
		want string
	}{
		iocs: []string{
			"example.com",
			"not.a.domain.example",
			"test.google.com",
		},
		want: ".*(example.com|not.a.domain.example|test.google.com)$",
	}
	if got := fmtRegex(tc.iocs); got != tc.want {
		t.Errorf("fmtRegex() returned unexpected regex string;\nwant: %s\ngot:%s\n", tc.want, got)
	}
}

func TestBadDomainFilter(t *testing.T) {
	tc := struct {
		bdd  *BadDomainDetection
		want []*nlpb.DNS
	}{
		bdd: &BadDomainDetection{
			logs: &NormalizedLog{
				dns: []*nlpb.DNS{
					{
						Query:  "1.2.3.4",
						Answer: "1.example.com",
					},
					{
						Query:  "4.3.2.1",
						Answer: "google.com",
					},
					{
						Query:  "domain.example",
						Answer: "not.a.domain.example",
					},
					{
						Query:  "test.google.com",
						Answer: "blah",
					},
				},
			},
			rr: regexp.MustCompile(".*example.com|not.a.domain.example|test.google.com$"),
		},
		want: []*nlpb.DNS{
			{
				Query:  "1.2.3.4",
				Answer: "1.example.com",
			},
			{
				Query:  "domain.example",
				Answer: "not.a.domain.example",
			},
			{
				Query:  "test.google.com",
				Answer: "blah",
			},
		},
	}
	got := tc.bdd.filter()
	if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
		t.Errorf("bdd.filter() returned unexpected diff (-want +got):\n%s", diff)
	}
}

func TestBadDomainAggregate(t *testing.T) {
	tc := struct {
		input []*nlpb.DNS
		want  []*spb.Signal
	}{
		input: []*nlpb.DNS{
			{
				Timestamp: &timestamppb.Timestamp{Seconds: 1718323200},
				SourceIp:  "1.2.3.4",
				Query:     "1.example.com",
				Answer:    "10.20.30.40",
			},
			{
				Timestamp: &timestamppb.Timestamp{Seconds: 1718323000},
				SourceIp:  "1.2.3.4",
				Query:     "domain.example",
				Answer:    "1.example.com",
			},
			{
				Timestamp: &timestamppb.Timestamp{Seconds: 1718322000},
				SourceIp:  "1.2.3.4",
				Query:     "2.example.com",
				Answer:    "40.30.20.10",
			},
			{
				Timestamp: &timestamppb.Timestamp{Seconds: 1718300000},
				SourceIp:  "4.3.2.1",
				Query:     "1.example.com",
				Answer:    "10.20.30.40",
			},
		},
		want: []*spb.Signal{
			{
				Event: &spb.Signal_BadDomain{
					BadDomain: &spb.BadDomain{
						TimestampStart: &timestamppb.Timestamp{Seconds: 1718322000},
						TimestampEnd:   &timestamppb.Timestamp{Seconds: 1718323200},
						SourceIp:       "1.2.3.4",
						BadDomain:      "example.com",
						DnsLog: []*nlpb.DNS{
							{
								Timestamp: &timestamppb.Timestamp{Seconds: 1718323200},
								SourceIp:  "1.2.3.4",
								Query:     "1.example.com",
								Answer:    "10.20.30.40",
							},
							{
								Timestamp: &timestamppb.Timestamp{Seconds: 1718323000},
								SourceIp:  "1.2.3.4",
								Query:     "domain.example",
								Answer:    "1.example.com",
							},
							{
								Timestamp: &timestamppb.Timestamp{Seconds: 1718322000},
								SourceIp:  "1.2.3.4",
								Query:     "2.example.com",
								Answer:    "40.30.20.10",
							},
						},
					},
				},
			},
			{
				Event: &spb.Signal_BadDomain{
					BadDomain: &spb.BadDomain{
						TimestampStart: &timestamppb.Timestamp{Seconds: 1718300000},
						TimestampEnd:   &timestamppb.Timestamp{Seconds: 1718300000},
						SourceIp:       "4.3.2.1",
						BadDomain:      "example.com",
						DnsLog: []*nlpb.DNS{
							{
								Timestamp: &timestamppb.Timestamp{Seconds: 1718300000},
								SourceIp:  "4.3.2.1",
								Query:     "1.example.com",
								Answer:    "10.20.30.40",
							},
						},
					},
				},
			},
		},
	}

	bdd := &BadDomainDetection{rr: regexp.MustCompile(".*(example.com|not.a.domain.example|test.google.com)$")}
	got := bdd.aggregate(tc.input)
	if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
		t.Errorf("bdd.aggregate() returned unexpected diff (-want +got):\n%s", diff)
	}
}
