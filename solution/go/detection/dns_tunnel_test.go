package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
)

func TestDNSTunnelFilter(t *testing.T) {
	tc := struct {
		dtd  *DNSTunnelDetection
		want []*nlpb.Netflow
	}{
		dtd: &DNSTunnelDetection{
			logs: &NormalizedLog{
				netflow: []*nlpb.Netflow{
					{
						SrcIp:   "1.2.3.4",
						SrcPort: 53,
						DstIp:   "4.3.2.1",
						DstPort: 59876,
					},
					{
						SrcIp:   "4.3.2.1",
						SrcPort: 59876,
						DstIp:   "1.2.3.4",
						DstPort: 53,
					},
					{
						SrcIp:   "4.3.2.1",
						SrcPort: 59876,
						DstIp:   "1.2.3.4",
						DstPort: 443,
					},
				},
			},
		},
		want: []*nlpb.Netflow{
			{
				SrcIp:   "1.2.3.4",
				SrcPort: 53,
				DstIp:   "4.3.2.1",
				DstPort: 59876,
			},
			{
				SrcIp:   "4.3.2.1",
				SrcPort: 59876,
				DstIp:   "1.2.3.4",
				DstPort: 53,
			},
		},
	}
	got := tc.dtd.filter()
	if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
		t.Errorf("dtd.filter() returned unexpected diff (-want +got):\n%s", diff)
	}
}

func TestDNSTunnelAggregate(t *testing.T) {
	tc := struct {
		input []*nlpb.Netflow
		want  []*spb.Signal
	}{
		input: []*nlpb.Netflow{
			{
				Timestamp: &timestamppb.Timestamp{Seconds: 1718321000},
				SrcIp:     "1.2.3.4",
				SrcPort:   53,
				DstIp:     "4.3.2.1",
				DstPort:   59876,
				BytesIn:   100,
				BytesOut:  0,
			},
			{
				Timestamp: &timestamppb.Timestamp{Seconds: 1718322000},
				SrcIp:     "4.3.2.1",
				SrcPort:   59876,
				DstIp:     "1.2.3.4",
				DstPort:   53,
				BytesOut:  400,
			},
			{
				Timestamp: &timestamppb.Timestamp{Seconds: 1718323000},
				SrcIp:     "4.3.2.1",
				SrcPort:   59876,
				DstIp:     "1.2.3.4",
				DstPort:   53,
				BytesOut:  400,
			},
			{
				Timestamp: &timestamppb.Timestamp{Seconds: 1718323200},
				SrcIp:     "9.8.7.6",
				SrcPort:   59876,
				DstIp:     "6.7.8.9",
				DstPort:   53,
				BytesIn:   0,
				BytesOut:  20,
			},
			{
				Timestamp: &timestamppb.Timestamp{Seconds: 1718323200},
				SrcIp:     "6.7.8.9",
				SrcPort:   53,
				DstIp:     "9.8.7.6",
				DstPort:   59876,
				BytesIn:   30,
				BytesOut:  0,
			},
		},
		want: []*spb.Signal{
			{
				Event: &spb.Signal_DnsTunnel{
					DnsTunnel: &spb.DNSTunnel{
						TimestampStart: &timestamppb.Timestamp{Seconds: 1718321000},
						TimestampEnd:   &timestamppb.Timestamp{Seconds: 1718323000},
						SourceIp:       "4.3.2.1",
						TunnelIp:       "1.2.3.4",
						BytesInTotal:   100,
						BytesOutTotal:  800,
						NetflowLog: []*nlpb.Netflow{
							{
								Timestamp: &timestamppb.Timestamp{Seconds: 1718321000},
								SrcIp:     "1.2.3.4",
								SrcPort:   53,
								DstIp:     "4.3.2.1",
								DstPort:   59876,
								BytesIn:   100,
								BytesOut:  0,
							},
							{
								Timestamp: &timestamppb.Timestamp{Seconds: 1718322000},
								SrcIp:     "4.3.2.1",
								SrcPort:   59876,
								DstIp:     "1.2.3.4",
								DstPort:   53,
								BytesOut:  400,
							},
							{
								Timestamp: &timestamppb.Timestamp{Seconds: 1718323000},
								SrcIp:     "4.3.2.1",
								SrcPort:   59876,
								DstIp:     "1.2.3.4",
								DstPort:   53,
								BytesOut:  400,
							},
						},
					},
				},
			},
		},
	}

	dtd := &DNSTunnelDetection{}
	got := dtd.aggregate(tc.input)
	if diff := cmp.Diff(tc.want, got, protocmp.Transform()); diff != "" {
		t.Errorf("dtd.aggregate() returned unexpected diff (-want +got):\n%s", diff)
	}
}
