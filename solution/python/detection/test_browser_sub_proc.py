from dataclasses import dataclass
from generated import normalized_log_pb2 as nlpb
from generated import signal_pb2 as spb
import browser_sub_proc
import helpers
import logging
import main
import unittest

class TestBrowserSubProcDetection(unittest.TestCase):
    def test_run(self):
        logging.disable(logging.WARNING)
        
        @dataclass
        class TestCase:
            desc: str
            bspb: browser_sub_proc.BrowserSubProcDetection
            want: list[spb.Signal]

        tests = [
            TestCase(
                desc = "Successful browser subprocess detection.",
                bspb = browser_sub_proc.BrowserSubProcDetection(
                    logs = main.NormalizedLog(
                        execution = [
                            # 001
                            #  | - 002 /usr/bin/Firefox &
                            #    | - 004 "/bin/bash blah"
                            #    | - 006 "/usr/bin/firefox blah"
                            #      | - 007 "/usr/bin/firefox blah blah"
                            #    | - 008 "/bin/bash /usr/bin/xdg-settings check default-web-browser firefox.desktop"
                            #  | - 003 "/usr/bin/curl blah | bash"
                            #    | - 005 "/bin/bash blah blah"
                            nlpb.Execution(
                                pid = 2,
                                ppid = 1,
                                filepath = "/usr/bin/Firefox",
                                command = "/usr/bin/Firefox &"
                            ),
                            nlpb.Execution(
                                pid = 3,
                                ppid = 1,
                                filepath = "/usr/bin/curl",
                                command = "/usr/bin/curl blah | bash"
                            ),
                            nlpb.Execution(
                                pid = 4,
                                ppid = 2,
                                filepath = "/bin/bash",
                                command = "/bin/bash blah",
                            ),
                            nlpb.Execution(
                                pid = 5,
                                ppid = 3,
                                filepath = "/bin/bash",
                                command = "/bin/bash blah blah"
                            ),
                            nlpb.Execution(
                                pid = 6,
                                ppid = 2,
                                filepath = "/usr/bin/firefox",
                                command = "/usr/bin/firefox blah"
                            ),
                            nlpb.Execution(
                                pid = 7,
                                ppid = 6,
                                filepath = "/usr/bin/firefox",
                                command = "/usr/bin/firefox blah blah"
                            ),
                            nlpb.Execution(
                                pid = 8,
                                ppid = 2,
                                filepath = "/bin/bash",
                                command = "/bin/bash /usr/bin/xdg-settings check default-web-browser firefox.desktop"
                            )
                        ]
                    )
                ),
                want = [
                    spb.Signal(
                        browser_sub_proc = spb.BrowserSubProc(
                            execution = nlpb.Execution(
                                pid = 4,
                                ppid = 2,
                                filepath = "/bin/bash",
                                command = "/bin/bash blah",
                            )
                        )
                    )
                ]
            )
        ]

        for test in tests:
            got = test.bspb.run()

            self.assertTrue(
                helpers.checkProtoListEqual(test.want, got)
            )

if __name__ == "__main__":
    unittest.main()