from dataclasses import dataclass
from generated import normalized_log_pb2 as nlpb
from google.protobuf.timestamp_pb2 import Timestamp
import helpers
import logging
import netflow
import unittest

class TestNetflowNormalizer(unittest.TestCase):
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
                desc = "Invalid source IP",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            ),
            TestCase(
                desc = "Invalid destination IP",
                # <TODO: Implement me!>
			    # input = ...
			    # want = ...
            ),
            TestCase(
			    desc = "Invalid source port",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            ),
            TestCase(
                desc ="Invalid destination port",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            ),
            TestCase(
                desc = "Invalid bytes in",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            ),
            TestCase(
                desc ="Invalid bytes out",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            ),
            TestCase(
                desc ="Invalid packets in",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            ),
            TestCase(
                desc ="Invalid packets out",
                # <TODO: Implement me!>
			    # input = ...,
			    # want = ...
            )
        ]
        for test in tests:
            nn = netflow.NetflowNormalizer("","")
            got = nn.normalize(test.input)
            try:
                if test.want is None:
                    self.assertEqual(test.want, got)
                else:
                    self.assertTrue(
                        helpers.checkProtoEqual(test.want, got)
                    )
            except AssertionError as e:
                logging.error("nn.normalize(%s) returned unexpected diff:\n%s", test.desc, e)
                raise e

if __name__ == "__main__":
    unittest.main()