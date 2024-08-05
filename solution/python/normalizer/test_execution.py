from dataclasses import dataclass
from generated import normalized_log_pb2 as nlpb
from google.protobuf.timestamp_pb2 import Timestamp
import execution
import helpers
import logging
import unittest

class TestExecutionNormalizer(unittest.TestCase):

    def test_normalize(self):
        logging.disable(logging.WARNING)

        @dataclass
        class TestCase:
            desc: str
            input: str
            want: nlpb.NormalizedLog

        tests = [
            TestCase(
                desc = "Successful validation",
                input = "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,4321,1234,\"/tmp\",\"hostname\",\"LINUX\"",
                want = nlpb.NormalizedLog(
                    execution_log=nlpb.Execution(
                        timestamp=Timestamp(seconds=1718323200),
                        filepath="/usr/bin/cat",
                        command="cat /tmp/secret",
                        uid=1000,
                        pid=4321,
                        ppid=1234,
                        cwd="/tmp",
                        hostname="hostname",
                        platform=nlpb.Execution.LINUX
                    )
                )
            ),
            TestCase(
                desc = "Invalid number of fields",
			    input = "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,4321,1234,\"/tmp\",\"hostname\"",
			    want = None
            ),
            TestCase(
                desc = "Invalid timestamp",
			    input = "2024-06-14 00:00:00,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,4321,1234,\"/tmp\",\"hostname\",\"LINUX\"",
			    want = None
            ),
            TestCase(
                desc = "Invalid UID",
			    input = "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",None,4321,1234,\"/tmp\",\"hostname\",\"LINUX\"",
			    want = None
            ),
            TestCase(
                desc = "Invalid PID",
			    input = "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,None,1234,\"/tmp\",\"hostname\",\"LINUX\"",
			    want = None,
            ),
            TestCase(
                desc = "Invalid PPID",
			    input = "1718323200,\"/usr/bin/cat\",\"cat /tmp/secret\",1000,4321,None,\"/tmp\",\"hostname\",\"LINUX\"",
			    want = None,
            )
        ]
        for test in tests:
            en = execution.ExecutionNormalizer("","")
            got = en.normalize(test.input)
            try:
                if test.want is None:
                    self.assertEqual(test.want, got)
                else:
                    self.assertTrue(
                        helpers.checkProtoEqual(test.want, got)
                    )
            except AssertionError as e:
                logging.error("en.normalize(%s) returned unexpected diff:\n%s", test.desc, e)
                raise e
            
if __name__ == "__main__":
    unittest.main()