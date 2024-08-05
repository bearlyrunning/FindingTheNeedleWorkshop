import argparse
import browser_sub_proc
import bad_domain
import dns_tunnel
import grpc
import logging
import os
from generated import enrichment_pb2_grpc as engrpc
from generated import signal_pb2 as spb
from google.protobuf.json_format import MessageToJson
from google.protobuf.json_format import Parse
from google.protobuf.json_format import ParseError

SIGNAL_PATH = "../../data/signal/{}.json"
BAD_DOMAIN_RULE_NAME = "bad_domain"
DNS_TUNNEL_RULE_NAME = "dns_tunnel"
BROWSER_SUB_PROC_RULE_NAME = "browser_sub_proc"

# Load signals to be enriched, into the given `enricher`.
def load(enricher) -> None:
        # Open signal file.
        try:
            f = open(enricher.signal_path.format(enricher.getName()), "r")
        except Exception as e:
            raise IOError("failed to open file %s, %s", enricher.signal_path, e)

        # Read signals and parse into protobufs.
        for s in f:    
            try:
                sig = Parse(s, spb.Signal())
            except Exception as e:
                raise ParseError("failed converting json signal to proto %s: %s", s, str(e))
            enricher.signals.append(sig)

        # Close input file.
        f.close()

# Output enriched signals to file.
def output(enricher) -> None:
    path = enricher.signal_path.format(enricher.getName() + "_enriched")

    # Check if the directory already exists, and make it otherwise.
    try:
        os.makedirs(os.path.dirname(path), exist_ok=True)
        out = open(path, "w+")
    except Exception as e:
        logging.error("failed to create output file %s: %s", path, str(e))
        raise e
    
        # Convert proto messages to JSON and write to file.
    for signal in enricher.signals:
        try:
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
    # Handle command line args
    parser = argparse.ArgumentParser(
        description="FindingTheNeedle's enrichment service client."
    )
    parser.add_argument('addr', default="localhost:8080", nargs="?", help="enrichment service address")
    flags = parser.parse_args()

    # Set up connection to server.
    if flags.addr == "localhost:8080":
        channel = grpc.insecure_channel(flags.addr)
    else:
        channel = grpc.secure_channel(
            flags.addr, 
            grpc.ssl_channel_credentials()
        )

    # Initiate a client
    c = engrpc.EnrichmentStub(channel)

    enrichers = [
        bad_domain.BadDomainEnricher(BAD_DOMAIN_RULE_NAME, SIGNAL_PATH),
        dns_tunnel.DNSTunnelEnricher(DNS_TUNNEL_RULE_NAME, SIGNAL_PATH),
        browser_sub_proc.BrowserSubProcEnricher(BROWSER_SUB_PROC_RULE_NAME, SIGNAL_PATH)
    ]
    
    # Load signals, enrich for each type, then output to file.
    for e in enrichers:
        try:
            load(e)
            e.enrich(c)
            output(e)
        except Exception as e:
            logging.fatal(e)

if __name__ == '__main__':
    main()