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
			    input = "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,90,10",
                want = nlpb.NormalizedLog(
                    netflow_log=nlpb.Netflow(
                        timestamp=Timestamp(seconds=1718323200),
                        log_source="test_logger_1",
                        src_ip="10.20.30.40",
                        src_port=123,
                        dst_ip="100.110.120.130",
                        dst_port=45678,
                        bytes_in=12345,
                        bytes_out=678,
                        packets_in=90,
                        packets_out=10,
                        protocol="UDP"
                    )
                )
            ),
            TestCase(
                desc = "Invalid number of fields",
			    input = "2024-06-14 00:00:00.000,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,90,10",
			    want = None
            ),
            TestCase(
                desc = "Invalid timestamp",
			    input = "2024-06-14 25:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,90,10",
			    want = None
            ),
            TestCase(
                desc = "Invalid source IP",
			    input = "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30,123,100.110.120.130,45678,12345,678,90,10",
			    want = None
            ),
            TestCase(
                desc = "Invalid destination IP",
			    input = "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120,45678,12345,678,90,10",
			    want = None
            ),
            TestCase(
			    desc = "Invalid source port",
			    input = "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,-1,100.110.120.130,45678,12345,678,90,10",
			    want = None
            ),
            TestCase(
                desc ="Invalid destination port",
			    input = "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,70000,12345,678,90,10",
			    want = None
            ),
            TestCase(
                desc = "Invalid bytes in",
			    input = "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,nil,678,90,10",
			    want = None
            ),
            TestCase(
                desc ="Invalid bytes out",
			    input = "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,nil,90,10",
			    want = None
            ),
            TestCase(
                desc ="Invalid packets in",
			    input = "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,nil,10",
			    want = None
            ),
            TestCase(
                desc ="Invalid packets out",
			    input = "2024-06-14 00:00:00.000,test_logger_1,UDP,10.20.30.40,123,100.110.120.130,45678,12345,678,90,nil",
			    want = None
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