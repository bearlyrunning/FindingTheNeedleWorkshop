package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
)

func TestNetflowNormalizer(t *testing.T) {
	var tests = []struct {
		desc string
		in   string
		want *nlpb.NormalizedLog
	}{
		{
			desc: "Successful validation",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid number of fields",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid timestamp",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid source IP",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid destination IP",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid source port",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid destination port",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid bytes in",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid bytes out",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid packets in",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid packets out",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
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
