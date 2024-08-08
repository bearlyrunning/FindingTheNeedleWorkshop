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
       	# <TODO: Implement me!>
	    # NOTE: This enrichment is slightly more complex compared to `enrichment/browser_sub_proc.py`
	    #       and `enrichment/dns_tunnel.py`. Please feel free to complete them in whatever order you prefer.
        # 
        # The bad_domain.csv file contains two fields:
	    #  - a domain indicator.
	    # - the OS this indicator is applicable to.
	    # Let's check if the source / client IP of the bad DNS traffic can be linked to a host with an applicable OS.
	    # - If the traffic is observed from an irrelevant OS, drop the signal.
	    # - If the host matches the relevant OS, populate the `hostname` field in the spb.BadDomain message for the signal.
	    # And then return the enriched signals.
        #
        # Hint #1: Check the RPC methods supported by enpb.EnrichmentClient in the generated package file.
        # Hint #2: What data structure can be used to allow easy lookup of the file content at `indicatorPath`?
        
        enrichedSigs = []

        for sig in self.signals:
            print("Implement me!")

        self.signals = enrichedSigs