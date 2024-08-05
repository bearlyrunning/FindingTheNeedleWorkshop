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
                # <TODO: Implement me!>
                # input = ...,
                # want = ...
            ),
            TestCase(
                desc = "Invalid number of fields",
			    # <TODO: Implement me!>
                # input = ...,
			    # want = ...
            ),
            TestCase(
                desc = "Invalid timestamp",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            ),
            TestCase(
                desc = "Invalid UID",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            ),
            TestCase(
                desc = "Invalid PID",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            ),
            TestCase(
                desc = "Invalid PPID",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
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