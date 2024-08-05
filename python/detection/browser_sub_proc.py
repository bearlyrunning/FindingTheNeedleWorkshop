import detection
from generated import signal_pb2 as spb

BROWSER_REGEX = ".*firefox$"
BROWSER_SUB_PROC_REGEX = ".*firefox.*"

class BrowserSubProcDetection(detection.DetectionInterface):
    def __init__(self, name="", logs=[]):
        self.name = name
        self.logs = logs

    # Get detection rule name.
    def ruleName(self) -> str:
        return self.name

    # Run detection logic.
    def run(self) -> list[spb.Signal]:
        # <TODO: Implement me!>
        # Find logs that represent suspicious browser child processes.
        #    1. Parse execution logs to surface the parent-child process relationship
	    #       (e.g. construct an execution tree using a set of lists and dicts).
	    #    2. Look for browser processes with subprocesses that are not browser-related subprocesses.
	    #       NOTE: to simplify the logic, if the commandline string matches browserSubProcRegex, it can be considered to be browser-related.
        #    3. Return the set of interesting logs as a list of spb.Signal.
        #
        # Expected output: the list of spb.Signal returned should have `browser_sub_proc` set.
        #
        # Hint #1: There are multiple approaches for surfacing the parent-child process relationship
	    #          (e.g. construct an execution tree using a set of lists and dicts), see example below:
        #          ```
        #              browserPids = []
        #              pids = {}           # dict[parent_pid]child_pid
        #              pidDetails = {}     # dict[pid]execution_log
        #              sigs = []
        #          ```
        # Hint #2: import re, use BROWSER_REGEX and BROWSER_SUB_PROC_REGEX
        # Hint #3: Check the fields you need to populate by inspecting the spb.BrowserSubProc protobuf message.
        return None