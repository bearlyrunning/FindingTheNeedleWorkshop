import enricher
import logging
from generated import enrichment_pb2 as enpb
from generated import signal_pb2 as spb

class BrowserSubProcEnricher(enricher.EnricherInterface):
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
        # For each signal, call our HostToIP enrichment.
        for sig in self.signals:
            try:
                resp = client.HostToIP(enpb.Host(
                    name=sig.browser_sub_proc.execution.hostname
                ))
                # Assign enriched IP address to protobuf field.
                sig.browser_sub_proc.source_ip = resp.ip
            except Exception as e:
                logging.warn("failed HostToIP conversion: %s", str(e))