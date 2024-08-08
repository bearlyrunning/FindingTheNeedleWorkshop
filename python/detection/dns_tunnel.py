import detection
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

    # Run detection logic.
    # run() function does the following:
    #   - loop through each proto in the /data/netflow/netflow_normalized.binpb or json file
    #     (these protos are saved in self.logs.netflow)
    #   - apply detection logic
    #   - output the log as a Signal proto (see /proto/signal.proto) message which
    #     eventually get saved in /data/signal/dns_tunnel.json
    def run(self) -> list[spb.Signal]:
        # <TODO: Implement me!>
        # Find relevant netflow logs indicating potential DNS tunneling behaviour.
        # To simplify the logic, the rule contains the following steps:
        #   1. Filter logs to what is relevant, then
        #   2. Aggregate logs based on source IP-destination IP address pairs.
        #   3. Only return logs with aggregated traffic volume above THRESHOLD.
        #   4. Return the set of interesting logs as a list of spb.Signal 
        #
        # Expected output: the list of spb.Signal returned should have `dns_tunnel` set.
        #
        # Hint #1: Assume DNS traffic is on port 53 inbound and outbound.
        # Hint #2: In your dict, construct a key using the values of source and destination IP addresses (for your pairing).
        # Hint #3: We need to compare traffic volume against a THRESHOLD - remember to keep a sum of bytes_in and bytes_out to check if we exceed the threshold.
        # Hint #3: Check the fields you need to populate by inspecting the spb.DNSTunnel protobuf message.
        
        sigs = []

        for log in self.logs.netflow:
            print("TODO: Implement me!")

        sigs.append(spb.Signal(
            dns_tunnel=spb.DNSTunnel(
                # timestamp_start=,
                # timestamp_end=,
                # source_ip=,
                # tunnel_ip=,
                # bytes_in_total=,
                # bytes_out_total=,
                # netflow_log=
            )
        ))

        return sigs