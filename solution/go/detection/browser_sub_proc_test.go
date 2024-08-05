package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
)

func TestBrowserSubProcDetection(t *testing.T) {
	bspb := &BrowserSubProcDetection{
		logs: &NormalizedLog{
			execution: []*nlpb.Execution{
				// 001
				//  | - 002 /usr/bin/Firefox &
				//    | - 004 "/bin/bash blah"
				//    | - 006 "/usr/bin/firefox blah"
				//      | - 007 "/usr/bin/firefox blah blah"
				//    | - 008 "/bin/bash /usr/bin/xdg-settings check default-web-browser firefox.desktop"
				//  | - 003 "/usr/bin/curl blah | bash"
				//    | - 005 "/bin/bash blah blah"
				{
					Pid:      int64(2),
					Ppid:     int64(1),
					Filepath: "/usr/bin/Firefox",
					Command:  "/usr/bin/Firefox &",
				},
				{
					Pid:      int64(3),
					Ppid:     int64(1),
					Filepath: "/usr/bin/curl",
					Command:  "/usr/bin/curl blah | bash",
				},
				{
					Pid:      int64(4),
					Ppid:     int64(2),
					Filepath: "/bin/bash",
					Command:  "/bin/bash blah",
				},
				{
					Pid:      int64(5),
					Ppid:     int64(3),
					Filepath: "/bin/bash",
					Command:  "/bin/bash blah blah",
				},
				{
					Pid:      int64(6),
					Ppid:     int64(2),
					Filepath: "/usr/bin/firefox",
					Command:  "/usr/bin/firefox blah",
				},
				{
					Pid:      int64(7),
					Ppid:     int64(6),
					Filepath: "/usr/bin/firefox",
					Command:  "/usr/bin/firefox blah blah",
				},
				{
					Pid:      int64(8),
					Ppid:     int64(2),
					Filepath: "/bin/bash",
					Command:  "/bin/bash /usr/bin/xdg-settings check default-web-browser firefox.desktop",
				},
			},
		},
	}

	want := []*spb.Signal{
		{
			Event: &spb.Signal_BrowserSubProc{
				BrowserSubProc: &spb.BrowserSubProc{
					Execution: &nlpb.Execution{
						Pid:      int64(004),
						Ppid:     int64(002),
						Filepath: "/bin/bash",
						Command:  "/bin/bash blah",
					},
				},
			},
		},
	}

	got, _ := bspb.run()
	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("bspb.run() returned unexpected diff (-want +got):\n%s", diff)
	}
}
