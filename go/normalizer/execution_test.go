package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/normalizedlogpb"
)

func TestSplitWithEscape(t *testing.T) {
	input := `1718366328,"/usr/bin/mount","/usr/bin/mount /var/lib/snapd/snaps/firefox_4483.snap /snap/firefox/4483 -t squashfs -o nodev\,ro\,x-gdu.hide\,x-gvfs-hide",0,10456,1,"/","bastion","LINUX"`
	want := []string{
		"1718366328",
		`"/usr/bin/mount"`,
		`"/usr/bin/mount /var/lib/snapd/snaps/firefox_4483.snap /snap/firefox/4483 -t squashfs -o nodev,ro,x-gdu.hide,x-gvfs-hide"`,
		"0",
		"10456",
		"1",
		`"/"`,
		`"bastion"`,
		`"LINUX"`,
	}

	got := splitWithEscape(input, ",", "\\")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("splitWithEscape(%s) returned unexpected diff (-want +got):\n%s", input, diff)
	}

}

func TestExecutionNormalizer(t *testing.T) {
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
			desc: "Invalid UID",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid PID",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
		{
			desc: "Invalid PPID",
			// <TODO: Implement me!>
			// in: ...,
			// want: ...,
		},
	}

	for _, tt := range tests {
		tt := tt // Added for legacy reason, prior to Go version 1.22, tt's memory location was reused within the loop.
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			en := &ExecutionNormalizer{}
			got := en.normalize(tt.in)
			if diff := cmp.Diff(tt.want, got, protocmp.Transform()); diff != "" {
				t.Errorf("en.normalize(%s) returned unexpected diff (-want +got):\n%s", tt.in, diff)
			}
		})
	}
}
