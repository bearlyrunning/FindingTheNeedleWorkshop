package main

import (
	"regexp"
	"strings"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
)

var (
	browserRegex        = regexp.MustCompile(".*firefox$")
	browserSubProcRegex = regexp.MustCompile(".*firefox.*")
)

func (bspd *BrowserSubProcDetection) ruleName() string {
	return bspd.name
}

func (bspd *BrowserSubProcDetection) run() ([]*spb.Signal, error) {
	var (
		browserPids []int64
		pids        = make(map[int64][]int64)         // map[parent_pid]child_pid
		pidDetails  = make(map[int64]*nlpb.Execution) // map[pid]execution_log
	)
	for _, e := range bspd.logs.execution {
		pids[e.GetPpid()] = append(pids[e.GetPpid()], e.GetPid())
		pidDetails[e.GetPid()] = e
		if browserRegex.MatchString(strings.ToLower(e.GetFilepath())) {
			browserPids = append(browserPids, e.GetPid())
		}
	}

	var sigs []*spb.Signal
	for _, bp := range browserPids {
		for _, pid := range pids[bp] {
			e := pidDetails[pid]
			// If a browser prcoess is spawning a non-browser-related subprocesses, output as anomalous.
			// NOTE: this is an oversimplified logic as we cannot entirely rely on commandline strings in real life.
			if !browserSubProcRegex.MatchString(strings.ToLower(e.GetCommand())) && e != nil {
				sigs = append(sigs, &spb.Signal{
					Event: &spb.Signal_BrowserSubProc{
						BrowserSubProc: &spb.BrowserSubProc{
							Execution: e,
						},
					},
				})
			}
		}
	}

	return sigs, nil
}
