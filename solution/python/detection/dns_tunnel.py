import detection
import google.protobuf.timestamp_pb2 as tspb
from generated import normalized_log_pb2 as nlpb
from generated import signal_pb2 as spb

# Threshold of total number of bytes for identifying DNS tunnelling.
THRESHOLD = 480

class DNSTunnelDetection(detection.DetectionInterface):
    def __init__(self, name="", logs=[]):
        self.name = name
        self.logs = logs

    # Get detection rule name.
    def ruleName(self) -> str:
        return self.name
    
    # Filter out netflow logs for DNS traffic.
    def filter(self) -> list[nlpb.Netflow]:
        matched = []

        # Look for logs related to DNS traffic (port 53).
        for n in self.logs.netflow:
            if n.dst_port == 53 or n.src_port == 53:
                matched.append(n)

        return matched

    # Aggregate traffic by IP pairs and count number of bytes transported over DNS.
    def aggregate(self, matched=[]) -> list[spb.Signal]:
        sigs = []
        keyByIpPair = {}
        for m in matched:
            # Set key format: client IP, remote IP
            k = m.src_ip + "," + m.dst_ip
            # If getting response from remote IP, reverse the order.
            if m.src_port == 53:
                k = m.dst_ip + "," + m.src_ip

            if k not in keyByIpPair:
                keyByIpPair[k] = [m]
            else:
                keyByIpPair[k].append(m)

        # Aggregate
        for pair, logs in keyByIpPair.items():
            earliest = logs[0].timestamp.ToDatetime()
            latest = logs[0].timestamp.ToDatetime()
            # Assuming bytes in and out numbers are set with directions of the traffic taken into consideration.
            bytes_in = logs[0].bytes_in
            bytes_out = logs[0].bytes_out

            for l in logs[1:]:
                bytes_in += l.bytes_in
                bytes_out += l.bytes_out
                if l.timestamp.ToDatetime() < earliest:
                    earliest = l.timestamp.ToDatetime()
                    continue
                if l.timestamp.ToDatetime() > latest:
                    latest = l.timestamp.ToDatetime()
            
            # Identify potential DNS tunnelling by checking total number of bytes against a static threshold.
            if bytes_out > THRESHOLD:
                ts_earliest = tspb.Timestamp()
                ts_earliest.FromDatetime(earliest)
                ts_latest = tspb.Timestamp()
                ts_latest.FromDatetime(latest)
                sigs.append(spb.Signal(
                    dns_tunnel=spb.DNSTunnel(
                        timestamp_start=ts_earliest,
                        timestamp_end=ts_latest,
                        source_ip=pair.split(',')[0],
                        tunnel_ip=pair.split(',')[1],
                        bytes_in_total=bytes_in,
                        bytes_out_total=bytes_out,
                        netflow_log=logs
                    )
                ))

        return sigs

    # Run detection logic.
    def run(self) -> list[spb.Signal]:
        return self.aggregate(self.filter())