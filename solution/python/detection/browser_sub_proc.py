import detection
import re
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
        browserPids = []
        pids = {}           # dict[parent_pid]child_pid
        pidDetails = {}     # dict[pid]execution_log
        sigs = []

        # Parse execution parent/child tree
        for e in self.logs.execution:
            if e.ppid not in pids:
                pids[e.ppid] = [e.pid]
            else:
                pids[e.ppid].append(e.pid)
            
            pidDetails[e.pid] = e
            if re.search(BROWSER_REGEX, e.filepath.lower()):
                browserPids.append(e.pid)

        # Populate signals
        for bp in browserPids:
            # Only look for browser processes with subprocesses
            if bp in pids:
                for pid in pids[bp]:
                    e = pidDetails[pid]
                    # If a browser prcoess is spawning a non-browser-related subprocesses, output as anomalous.
                    # NOTE: this is an oversimplified logic as we cannot entirely rely on commandline strings in real life.
                    if not re.search(BROWSER_SUB_PROC_REGEX, e.command.lower()) and e is not None:
                        sigs.append(spb.Signal(
                            browser_sub_proc=spb.BrowserSubProc(
                                execution = e
                            )
                        ))
        
        return sigs