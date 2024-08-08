package main

import (
	"fmt"
	"regexp"

	spb "github.com/bearlyrunning/FindingTheNeedle/go/generated/signalpb"
)

var (
	browserRegex        = regexp.MustCompile(".*firefox$")
	browserSubProcRegex = regexp.MustCompile(".*firefox.*")
)

func (bspd *BrowserSubProcDetection) ruleName() string {
	return bspd.name
}

// run() function does the following:
//   - loop through each proto in the /data/execution/execution_normalized.binpb or json file
//     (these protos are saved in bdd.logs.execution)
//   - apply detection logic
//   - output the log as a Signal proto (see /proto/signal.proto) message which
//     eventually get saved in /data/signal/browser_sub_proc.json
func (bspd *BrowserSubProcDetection) run() ([]*spb.Signal, error) {
	// <TODO: Implement me!>
	// Find logs that represent suspicious browser child processes.
	//    1. Parse execution logs to surface the parent-child process relationship
	//       (e.g. construct an execution tree using a set of slices and maps).
	//    2. Look for browser processes with subprocesses that are not browser-related subprocesses.
	//       NOTE: to simplify the logic, if the commandline string matches browserSubProcRegex, it can be considered to be browser-related.
	//    3. Return the set of interesting logs as a list of spb.Signal.

	// Expected output: the list of spb.Signal returned should have `event` field set to `browser_sub_proc`.

	// Hint #1: There are multiple approaches for surfacing the parent-child process relationship
	//          (e.g. construct an execution tree using a set of slices and maps), see example below:
	//          ```
	//          var (
	//          	  browserPids []int64
	//          	  pids        = make(map[int64][]int64)         // map[parent_pid]child_pid
	//          	  pidDetails  = make(map[int64]*nlpb.Execution) // map[pid]execution_log
	//          	  sigs        []*spb.Signal
	//          )
	// ```
	// Hint #2: Make use of `regexp` package, browserRegex and browserSubProcRegex
	// Hint #3: Check the fields you need to populate by inspecting the spb.BrowserSubProc protobuf message.

	var sigs []*spb.Signal

	for _, e := range bspd.logs.execution {
		fmt.Printf("TODO: Implement me! %v", e)
	}

	sigs = append(sigs, &spb.Signal{
		Event: &spb.Signal_BrowserSubProc{
			BrowserSubProc: &spb.BrowserSubProc{
				// Execution: e,
			},
		},
	})

	return sigs, nil
}
