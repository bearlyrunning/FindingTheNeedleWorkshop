import enricher
import logging
from generated import enrichment_pb2 as enpb
from generated import signal_pb2 as spb

class DNSTunnelEnricher(enricher.EnricherInterface):
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
	    # Populate the `hostname` field in the spb.DNSTunnel message for all signals.
	    # Hint: Check the RPC methods supported by enpb.EnrichmentClient in the generated package file.

        # For each signal, call our IPToHost enrichment.
        for sig in self.signals:
            try:
                print("Implement me!")
            except Exception as e:
                logging.warn("failed IPToHost conversion: %s", str(e))