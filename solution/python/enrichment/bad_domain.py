import enricher
import logging
from generated import enrichment_pb2 as enpb
from generated import signal_pb2 as spb
from typing import Dict


INDICATOR_PATH = "../../data/indicators/bad_domain.csv"
COLUMN_NUM     = 2

def load_indicators() -> Dict[str, enpb.Host.Platform]:
    hostToOS = {} # dict[string]enpb.Host_Platform
    # Open indicators
    try:
        lst = open(INDICATOR_PATH)
    except Exception as e:
            logging.error("failed to open file %s: %s", INDICATOR_PATH, str(e))
            raise e

    for i in lst:
        strs = i.strip().split(",")
        hostToOS[strs[0]] = enpb.Host.Platform.Value(strs[1])

    lst.close()
    return hostToOS

class BadDomainEnricher(enricher.EnricherInterface):
    def __init__(self, name="", path=""):
        self.signals = []
        self.name = name
        self.signal_path = path

    # Get enricher name.
    def getName(self) -> str:
        return self.name
    
    # Get signals associated with a given enricher.
    def getSignals(self) -> list[spb.Signal]:
        return self.signals
    
    # Set the signals to be enriched.
    def setSignals(self, signals=[]) -> None:
        self.signals = signals
    
    # Enrich signals based on the provided client.
    def enrich(self, client) -> None:
        enrichedSigs = []
        hostOSDict = load_indicators()
        # For each signal, call our IPToHost enrichment based on source IP address.
        for sig in self.signals:
            try:
                resp = client.IPToHost(enpb.IP(
                    ip=sig.bad_domain.source_ip
                ))
                if resp.platform != hostOSDict[sig.bad_domain.bad_domain]:
                    logging.warn("domain indicator %s observed from %s - an operating system that isn't impacted", sig.bad_domain.bad_domain, enpb.Host.Platform.Name(resp.platform))
                else:
                    # Assign enriched hostname to protobuf field.
                    sig.bad_domain.hostname = resp.name
                    enrichedSigs.append(sig)
            except Exception as e:
                logging.warn("failed IPToHost conversion, output signals for manual investigation: %s", str(e))
                enrichedSigs.append(sig)
        self.signals = enrichedSigs