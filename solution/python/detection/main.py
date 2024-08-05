import bad_domain
import browser_sub_proc
import dns_tunnel
import logging
import os
import struct
import sys
from generated import normalized_log_pb2 as nlpb
from google.protobuf.json_format import MessageToJson

BAD_DOMAIN_RULE_NAME = "bad_domain"
DNS_TUNNEL_RULE_NAME = "dns_tunnel"
BROWSER_SUB_PROC_FULE_NAME = "browser_sub_proc"
LOG_PATHS = [
    "../../data/dns/dns_normalized.binpb",
    "../../data/netflow/netflow_normalized.binpb",
    "../../data/execution/execution_normalized.binpb"
]
SIGNAL_PATH = "../../data/signal/{}.json"
SIZE_BUF = 4

class NormalizedLog:
    def __init__(self, dns=[], netflow=[], execution=[]):
        self.dns = dns
        self.netflow = netflow
        self.execution = execution

    # Load normalized log data.
    def load(self) -> None:
        for path in LOG_PATHS:
            try:
                file = open(path, 'rb')
            except Exception as e:
                logging.error("failed to open file %s: %s", path, e)

            filesize = os.path.getsize(path)
            # Iterate through file reading size:protobuf pairs.
            while file.tell() < filesize:
                try:
                    # Read the prepending proto length
                    size = file.read(SIZE_BUF)
                    protosize = struct.unpack('<i', size)[0]
                except Exception as e:
                    logging.error("error parsing file: %s", str(e))
                    raise

                # Read serialized proto from file.
                msgBuf = file.read(protosize)

                msg = nlpb.NormalizedLog()
                # Parse into a proto object.
                try:    
                    msg.ParseFromString(msgBuf)
                except Exception as e:
                    logging.error("failed decoding log message %s: %s", msg, e)

                # Determine which log the read proto relates to, and append to relevant list.
                if msg.HasField("dns_log"):
                    self.dns.append(msg.dns_log)
                elif msg.HasField("netflow_log"):
                    self.netflow.append(msg.netflow_log)
                elif msg.HasField("execution_log"):
                    self.execution.append(msg.execution_log)

            file.close()
        return 

# Output signals to a file of given name.
def output(name, sigs) -> None:
# Set up output file for signals.
    path = SIGNAL_PATH.format(name)
    try:
        # Check if the directory already exists, and make it otherwise.
        os.makedirs(os.path.dirname(path), exist_ok=True)
        out = open(path, "w+")
    except Exception as e:
        logging.error("failed to create output file %s: %s", path, str(e))
        raise e

    # Convert proto messages to JSON and write to file.
    for signal in sigs:
        try:
            # Write protobuf as JSON on a single line.
            j = MessageToJson(signal).replace('\n','').replace(' ', '')
        except Exception as e:
            logging.error("failed converting log message to json format %s: %s", signal, str(e))
            raise e
        out.write(j)
        out.write('\n')

    # Close output file.
    out.close()

# Entry point.
def main() -> None:
    # Read normalized logs in from disk.
    nl = NormalizedLog()
    try:
        nl.load()
    except Exception as e:
        logging.fatal("failed loading normalized logs: %s", e)
        return None
    
    # Run detection logic.
    detections = [
        bad_domain.BadDomainDetection(
            name=BAD_DOMAIN_RULE_NAME,
            logs=nl,
            indicators=[],
        ),
        dns_tunnel.DNSTunnelDetection(
            name=DNS_TUNNEL_RULE_NAME,
            logs=nl
        ),
        browser_sub_proc.BrowserSubProcDetection(
            name=BROWSER_SUB_PROC_FULE_NAME,
            logs=nl
        )
    ]

    for d in detections:
        try:
            sigs = d.run()
        except Exception as e:
            logging.fatal("failed to run detection rule %s: %s", d.name, str(e))
            sys.exit(1)

        output(d.ruleName(), sigs)

if __name__ == '__main__':
    main()