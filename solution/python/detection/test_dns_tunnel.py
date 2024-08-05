from dataclasses import dataclass
from generated import normalized_log_pb2 as nlpb
from generated import signal_pb2 as spb
import google.protobuf.timestamp_pb2 as tspb
import dns_tunnel
import helpers
import logging
import main
import unittest

class TestDNSTunnel(unittest.TestCase):
    def test_filter(self):
        logging.disable(logging.WARNING)
        
        @dataclass
        class TestCase:
            desc: str
            dtd: dns_tunnel.DNSTunnelDetection
            want: list[nlpb.Netflow]

        tests = [
            TestCase(
                desc = "Successful aggregation.",
                dtd = dns_tunnel.DNSTunnelDetection(
                    logs = main.NormalizedLog(
                        netflow = [
                            nlpb.Netflow(
                                src_ip = "1.2.3.4",
                                src_port = 53,
                                dst_ip = "4.3.2.1",
                                dst_port = 59876
                            ),
                            nlpb.Netflow(
                                src_ip = "4.3.2.1",
                                src_port = 59876,
                                dst_ip = "1.2.3.4",
                                dst_port = 53
                            ),
                            nlpb.Netflow(
                                src_ip = "4.3.2.1",
                                src_port = 59876,
                                dst_ip = "1.2.3.4",
                                dst_port = 443
                            ),
                        ]
                    )
                ),
                want = [
                    nlpb.Netflow(
                        src_ip = "1.2.3.4",
                        src_port = 53,
                        dst_ip = "4.3.2.1",
                        dst_port = 59876
                    ),
                    nlpb.Netflow(
                        src_ip = "4.3.2.1",
                        src_port = 59876,
                        dst_ip = "1.2.3.4",
                        dst_port = 53
                    ),
                ]
            )
        ]

        for test in tests:
            got = test.dtd.filter()
            self.assertTrue(
                helpers.checkProtoListEqual(test.want, got)
            )
    
    def test_aggregate(self):
        logging.disable(logging.WARNING)

        @dataclass
        class TestCase:
            desc: str
            input: list[nlpb.Netflow]
            want: list[spb.Signal]

        tests = [
            TestCase(
                desc = "Successful aggregation.",
                input = [
                    nlpb.Netflow(
                        timestamp = tspb.Timestamp(seconds=1718321000),
                        src_ip = "1.2.3.4",
                        src_port = 53,
                        dst_ip = "4.3.2.1",
                        dst_port = 59876,
                        bytes_in = 100,
                        bytes_out = 0
                    ),
                    nlpb.Netflow(
                        timestamp = tspb.Timestamp(seconds=1718322000),
                        src_ip = "4.3.2.1",
                        src_port = 59876,
                        dst_ip = "1.2.3.4",
                        dst_port = 53,
                        bytes_out = 400
                    ),
                    nlpb.Netflow(
                        timestamp = tspb.Timestamp(seconds=1718323000),
                        src_ip = "4.3.2.1",
                        src_port = 59876,
                        dst_ip = "1.2.3.4",
                        dst_port = 53,
                        bytes_out = 400
                    ),
                    nlpb.Netflow(
                        timestamp = tspb.Timestamp(seconds=1718323200),
                        src_ip = "9.8.7.6",
                        src_port = 59876,
                        dst_ip = "6.7.8.9",
                        dst_port = 53,
                        bytes_in = 0,
                        bytes_out = 20
                    ),
                    nlpb.Netflow(
                        timestamp = tspb.Timestamp(seconds=1718323200),
                        src_ip = "6.7.8.9",
                        src_port = 53,
                        dst_ip = "9.8.7.6",
                        dst_port = 59876,
                        bytes_in = 30,
                        bytes_out = 0
                    )
                ],
                want = [
                    spb.Signal(
                        dns_tunnel = spb.DNSTunnel(
                            timestamp_start = tspb.Timestamp(seconds=1718321000),
                            timestamp_end = tspb.Timestamp(seconds=1718323000),
                            source_ip = "4.3.2.1",
                            tunnel_ip = "1.2.3.4",
                            bytes_in_total = 100,
                            bytes_out_total = 800,
                            netflow_log = [
                                nlpb.Netflow(
                                    timestamp = tspb.Timestamp(seconds=1718321000),
                                    src_ip = "1.2.3.4",
                                    src_port = 53,
                                    dst_ip = "4.3.2.1",
                                    dst_port = 59876,
                                    bytes_in = 100,
                                    bytes_out = 0
                                ),
                                nlpb.Netflow(
                                    timestamp = tspb.Timestamp(seconds=1718322000),
                                    src_ip = "4.3.2.1",
                                    src_port = 59876,
                                    dst_ip = "1.2.3.4",
                                    dst_port = 53,
                                    bytes_out = 400
                                ),
                                nlpb.Netflow(
                                    timestamp = tspb.Timestamp(seconds=1718323000),
                                    src_ip = "4.3.2.1",
                                    src_port = 59876,
                                    dst_ip = "1.2.3.4",
                                    dst_port = 53,
                                    bytes_out = 400
                                )
                            ]
                        )
                    )
                ]   
            )
        ]

        for test in tests:
            dtd = dns_tunnel.DNSTunnelDetection()
            got = dtd.aggregate(test.input)

            self.assertTrue(
                helpers.checkProtoListEqual(test.want,got)
            )

if __name__ == "__main__":
    unittest.main()