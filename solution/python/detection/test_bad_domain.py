from dataclasses import dataclass
from generated import normalized_log_pb2 as nlpb
from generated import signal_pb2 as spb
import google.protobuf.timestamp_pb2 as tspb
import bad_domain
import helpers
import logging
import main
import unittest

class TestBadDomain(unittest.TestCase):
    def test_fmt_regex(self):
        logging.disable(logging.WARNING)
        
        @dataclass
        class TestCase:
            desc: str
            iocs: list
            want: str
        
        tests = [
            TestCase(
                desc = "Successful regex formatting.",
                iocs = [
                    "example.com",
			        "not.a.domain.example",
			        "test.google.com",
                ],
                want = ".*(example.com|not.a.domain.example|test.google.com)$"
            )
        ]
        for test in tests:
            bdr = bad_domain.BadDomainDetection(indicators=test.iocs)
            bdr.fmtRegex()
            self.assertEqual(bdr.regStr, test.want)

    def test_filter(self):
        logging.disable(logging.WARNING)
        
        @dataclass
        class TestCase:
            desc: str
            bdr: bad_domain.BadDomainDetection
            want: list[nlpb.DNS]

        tests = [
            TestCase(
                desc = "Successful indicator match",
                bdr = bad_domain.BadDomainDetection(
                    logs = main.NormalizedLog(
                        dns = [
                            nlpb.DNS(
                                query = "1.2.3.4",
                                answer = "1.example.com"
                            ),
                            nlpb.DNS(
                                query = "4.3.2.1",
                                answer = "google.com"
                            ),
                            nlpb.DNS(
                                query = "domain.example",
                                answer = "not.a.domain.example"
                            ),
                            nlpb.DNS(
                                query = "test.google.com",
                                answer = "blah"
                            )
                        ]
                    ), 
                    indicators = ["example.com", "not.a.domain.example", "test.google.com"],
                ),
                want = [
                    nlpb.DNS(
                        query = "1.2.3.4",
                        answer = "1.example.com"
                    ),
                    nlpb.DNS(
                        query = "domain.example",
                        answer = "not.a.domain.example"
                    ),
                    nlpb.DNS(
                        query = "test.google.com",
                        answer = "blah"
                    )
                ]
            )
        ]

        for test in tests:
            got = test.bdr.filter()
            self.assertTrue(
                helpers.checkProtoListEqual(test.want,got)
            )

    def test_aggregate(self):
        logging.disable(logging.WARNING)
        
        @dataclass
        class TestCase:
            desc: str
            input: list[nlpb.DNS]
            want: list[spb.Signal]

        tests = [
            TestCase(
                desc = "Successful aggregation.",
                input = [
                    nlpb.DNS(
                        timestamp = tspb.Timestamp(seconds=1718323200),
                        source_ip = "1.2.3.4",
                        query = "1.example.com",
                        answer = "10.20.30.40",
                    ),
                    nlpb.DNS(
                        timestamp = tspb.Timestamp(seconds=1718323000),
                        source_ip = "1.2.3.4",
                        query = "domain.example",
                        answer = "1.example.com",
                    ),
                    nlpb.DNS(
                        timestamp = tspb.Timestamp(seconds=1718322000),
                        source_ip = "1.2.3.4",
                        query = "2.example.com",
                        answer = "40.30.20.10",
                    ),
                    nlpb.DNS(
                        timestamp = tspb.Timestamp(seconds=1718300000),
                        source_ip = "4.3.2.1",
                        query = "1.example.com",
                        answer = "10.20.30.40",
                    ),
                ],
                want = [
                    spb.Signal(
                        bad_domain=spb.BadDomain(
                            timestamp_start = tspb.Timestamp(seconds=1718322000),
                            timestamp_end = tspb.Timestamp(seconds=1718323200),
                            source_ip = "1.2.3.4",
                            bad_domain = "example.com",
                            dns_log=[
                                nlpb.DNS(
                                    timestamp = tspb.Timestamp(seconds=1718323200),
                                    source_ip = "1.2.3.4",
                                    query = "1.example.com",
                                    answer = "10.20.30.40",
                                ),
                                nlpb.DNS(
                                    timestamp = tspb.Timestamp(seconds=1718323000),
                                    source_ip = "1.2.3.4",
                                    query = "domain.example",
                                    answer = "1.example.com",
                                ),
                                nlpb.DNS(
                                    timestamp = tspb.Timestamp(seconds=1718322000),
                                    source_ip = "1.2.3.4",
                                    query = "2.example.com",
                                    answer = "40.30.20.10",
                                ),
                            ]
                        )
                    ),
                    spb.Signal(
                        bad_domain=spb.BadDomain(
                            timestamp_start = tspb.Timestamp(seconds=1718300000),
                            timestamp_end = tspb.Timestamp(seconds=1718300000),
                            source_ip = "4.3.2.1",
                            bad_domain = "example.com",
                            dns_log=[
                                nlpb.DNS(
                                    timestamp = tspb.Timestamp(seconds=1718300000),
                                    source_ip = "4.3.2.1",
                                    query = "1.example.com",
                                    answer = "10.20.30.40",
                                )
                            ]
                        )
                    )
                ]
            )
        ]

        for test in tests:
            bdd = bad_domain.BadDomainDetection(
                indicators = ["example.com", "not.a.domain.example", "test.google.com"],
            )
            bdd.fmtRegex()
            got = bdd.aggregate(test.input)
            
            self.assertTrue(
                helpers.checkProtoListEqual(test.want,got)
            )

if __name__ == "__main__":
    unittest.main()