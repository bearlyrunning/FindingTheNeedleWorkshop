#!/usr/bin/env python3
import dns
import execution
import netflow
import sys
from google.protobuf.json_format import MessageToJson

DNS_LOG     = "../../data/dns/dns.log"
DNS_BINARY_OUTPUT  = "../../data/dns/dns_normalized.binpb"
DNS_JSON_OUTPUT  = "../../data/dns/dns_normalized.json"
NET_LOG     = "../../data/netflow/netflow.log"
NET_BINARY_OUTPUT  = "../../data/netflow/netflow_normalized.binpb"
NET_JSON_OUTPUT  = "../../data/netflow/netflow_normalized.json"
EXEC_LOG    = "../../data/execution/execution.log"
EXEC_BINARY_OUTPUT = "../../data/execution/execution_normalized.binpb"
EXEC_JSON_OUTPUT = "../../data/execution/execution_normalized.json"

# Run normalizer logic.
def run(normalizer) -> None:
    # Open file and set up reader
    reader = open(normalizer.getInput(), "r")

    # Set up wire / binary output writer.
    writer_bin = open(normalizer.getBinaryOutput(), "wb")

    # Set up JSON output writer.
    writer_json = open(normalizer.getJsonOutput(), "w")

    # Normalize line by line.
    for line in reader:
        msg = normalizer.normalize(line)

        if msg is not None:
            en = msg.SerializeToString()

            # Write the normalized protobuf, prepended with size as a delimiter.
            # See https://protobuf.dev/programming-guides/techniques/#streaming.
            writer_bin.write(prependSize(en))

            writer_json.write(MessageToJson(msg)+"\n")

    # Close files
    reader.close()
    writer_bin.close()
    writer_json.close()

# Prepend binary size to given proto.
def prependSize(proto) -> bytes:
    return len(proto).to_bytes(4 ,'little') + proto

# Entry point.
def main() -> None:
    normalizers = [
        dns.DNSNormalizer(DNS_LOG, DNS_BINARY_OUTPUT, DNS_JSON_OUTPUT),
        netflow.NetflowNormalizer(NET_LOG, NET_BINARY_OUTPUT, NET_JSON_OUTPUT),
        execution.ExecutionNormalizer(EXEC_LOG, EXEC_BINARY_OUTPUT, EXEC_JSON_OUTPUT)
    ]
    for n in normalizers:
        try:
            run(n)
        except Exception as e:
            print(str(e))
            sys.exit(1)

if __name__ == '__main__':
    main()