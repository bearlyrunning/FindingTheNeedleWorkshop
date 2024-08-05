package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
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
			in:   "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,4321,1234,\"/tmp\",\"hostname\",\"LINUX\"",
			want: &nlpb.NormalizedLog{
				Msg: &nlpb.NormalizedLog_ExecutionLog{
					ExecutionLog: &nlpb.Execution{
						Timestamp: &timestamppb.Timestamp{Seconds: 1718323200},
						Filepath:  "/usr/bin/cat",
						Command:   "cat /tmp/secret",
						Uid:       int64(1000),
						Pid:       int64(4321),
						Ppid:      int64(1234),
						Cwd:       "/tmp",
						Hostname:  "hostname",
						Platform:  nlpb.Execution_LINUX,
					},
				},
			},
		},
		{
			desc: "Invalid number of fields",
			in:   "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,4321,1234,\"/tmp\",\"hostname\"",
			want: nil,
		},
		{
			desc: "Invalid timestamp",
			in:   "2024-06-14 00:00:00,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,4321,1234,\"/tmp\",\"hostname\",\"LINUX\"",
			want: nil,
		},
		{
			desc: "Invalid UID",
			in:   "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",None,4321,1234,\"/tmp\",\"hostname\",\"LINUX\"",
			want: nil,
		},
		{
			desc: "Invalid PID",
			in:   "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,None,1234,\"/tmp\",\"hostname\",\"LINUX\"",
			want: nil,
		},
		{
			desc: "Invalid PPID",
			in:   "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,4321,None,\"/tmp\",\"hostname\",\"LINUX\"",
			want: nil,
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
