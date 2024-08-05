from dataclasses import dataclass
from generated import normalized_log_pb2 as nlpb
from google.protobuf.timestamp_pb2 import Timestamp
import dns
import helpers
import logging
import unittest

class TestDNSNormalizer(unittest.TestCase):
    
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
                input = "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14,0",
                want = nlpb.NormalizedLog(
                    dns_log=nlpb.DNS(
                        timestamp=Timestamp(seconds=1718323200),
                        log_source="test_logger_1",
                        source_ip="10.20.30.40",
                        resolver_ip="100.110.120.130",
                        query="example.com",
                        type="A",
                        answer="93.184.215.14",
                        return_code=nlpb.DNS.NOERROR
                    )
                ),
            ),
            TestCase(
                desc = "Invalid number of fields",
                input = "2024-06-14 25:00,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14",
                want = None
            ),
            TestCase(
                desc = "Invalid timestamp",
			    input = "2024-06-14 25:00,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14,0",
			    want = None
            ),
            TestCase(
                desc = "Invalid source IP",
			    input = "2024-06-14 00:00:00.000,test_logger_1,10.20.30,100.110.120.130,example.com,A,93.184.215.14,0",
			    want = None
            ),
            TestCase(
                desc = "Invalid DNS resolver IP",
			    input = "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120,example.com,A,93.184.215.14,0",
			    want = None
            ),
             TestCase(
                desc = "Invalid query",
			    input = "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120.130,,A,93.184.215.14,0",
			    want = None
            ),
            TestCase(
                desc = "Invalid return code non integer",
			    input = "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14,none",
			    want = None
            ),
            TestCase(
                desc = "Invalid return code",
			    input = "2024-06-14 00:00:00.000,test_logger_1,10.20.30.40,100.110.120.130,example.com,A,93.184.215.14,10",
			    want = None
            )
        ]
        for test in tests:
            dn = dns.DNSNormalizer("","")
            got = dn.normalize(test.input)
            try:
                if test.want is None:
                    self.assertEqual(test.want, got)
                else:
                    self.assertTrue(
                        helpers.checkProtoEqual(test.want, got)
                    )
            except AssertionError as e:
                logging.error("dn.normalize(%s) returned unexpected diff:\n%s", test.desc, e)
                raise e

if __name__ == "__main__":
    unittest.main()