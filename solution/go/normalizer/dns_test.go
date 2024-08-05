package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
)

func TestDNSNormalizer(t *testing.T) {
	var tests = []struct {
		desc string
		in   string
		want *nlpb.NormalizedLog
	}{
		{
			desc: "Successful validation",
			in:   "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14,0",
			want: &nlpb.NormalizedLog{
				Msg: &nlpb.NormalizedLog_DnsLog{
					DnsLog: &nlpb.DNS{
						Timestamp:  &timestamppb.Timestamp{Seconds: 1718323200},
						Query:      "example.com",
						Type:       "A",
						Answer:     "93.184.215.14",
						ReturnCode: nlpb.DNS_NOERROR,
						SourceIp:   "10.20.30.40",
						ResolverIp: "100.110.120.130",
						LogSource:  "test_logger_1",
					},
				},
			},
		},
		{
			desc: "Invalid number of fields",
			in:   "2024-06-14 25:00,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14",
			want: nil,
		},
		{
			desc: "Invalid timestamp",
			in:   "2024-06-14 25:00,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14,0",
			want: nil,
		},
		{
			desc: "Invalid source IP",
			in:   "2024-06-14 00:00:00.000,test_logger_1,10.20.30,100.110.120.130,example.com,A,93.184.215.14,0",
			want: nil,
		},
		{
			desc: "Invalid DNS resolver IP",
			in:   "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120,example.com,A,93.184.215.14,0",
			want: nil,
		},
		{
			desc: "Invalid query",
			in:   "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120.130,,A,93.184.215.14,0",
			want: nil,
		},
		{
			desc: "Invalid return code non integer",
			in:   "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14,none",
			want: nil,
		},
		{
			desc: "Invalid return code",
			in:   "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14,10",
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt // Added for legacy reason, prior to Go version 1.22, tt's memory location was reused within the loop.
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			dn := &DNSNormalizer{}
			got := dn.normalize(tt.in)
			if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("dn.normalize(%s) returned unexpected diff (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
