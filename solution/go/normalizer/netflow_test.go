package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
)

func TestNetflowNormalizer(t *testing.T) {
	var tests = []struct {
		desc string
		in   string
		want *nlpb.NormalizedLog
	}{
		{
			desc: "Successful validation",
			in:   "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,90,10",
			want: &nlpb.NormalizedLog{
				Msg: &nlpb.NormalizedLog_NetflowLog{
					NetflowLog: &nlpb.Netflow{
						Timestamp:  &timestamppb.Timestamp{Seconds: 1718323200},
						SrcIp:      "10.20.30.40",
						SrcPort:    int32(123),
						DstIp:      "100.110.120.130",
						DstPort:    int32(45678),
						BytesIn:    int64(12345),
						BytesOut:   int64(678),
						PacketsIn:  int64(90),
						PacketsOut: int64(10),
						Protocol:   "UDP",
						LogSource:  "test_logger_1",
					},
				},
			},
		},
		{
			desc: "Invalid number of fields",
			in:   "2024-06-14 00:00:00.000,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,90,10",
			want: nil,
		},
		{
			desc: "Invalid timestamp",
			in:   "2024-06-14 25:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,90,10",
			want: nil,
		},
		{
			desc: "Invalid source IP",
			in:   "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30,123,100.110.120.130,45678,12345,678,90,10",
			want: nil,
		},
		{
			desc: "Invalid destination IP",
			in:   "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120,45678,12345,678,90,10",
			want: nil,
		},
		{
			desc: "Invalid source port",
			in:   "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,-1,100.110.120.130,45678,12345,678,90,10",
			want: nil,
		},
		{
			desc: "Invalid destination port",
			in:   "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,70000,12345,678,90,10",
			want: nil,
		},
		{
			desc: "Invalid bytes in",
			in:   "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,nil,678,90,10",
			want: nil,
		},
		{
			desc: "Invalid bytes out",
			in:   "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,nil,90,10",
			want: nil,
		},
		{
			desc: "Invalid packets in",
			in:   "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,nil,10",
			want: nil,
		},
		{
			desc: "Invalid packets out",
			in:   "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,90,nil",
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt // Added for legacy reason, prior to Go version 1.22, tt's memory location was reused within the loop.
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			nn := &NetflowNormalizer{}
			got := nn.normalize(tt.in)
			if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("nn.normalize(%s) returned unexpected diff (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
