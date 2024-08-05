import detection
import logging
from generated import signal_pb2 as spb

INDICATOR_PATH = "../../data/indicators/bad_domain.csv"

class BadDomainDetection(detection.DetectionInterface):
    def __init__(self, name="", logs=[], indicators=[], signals=[]):
        self.name = name
        self.logs = logs
        self.indicators = indicators
        self.signals = signals
        self.regStr = ""

    # Get detection rule name.
    def ruleName(self) -> str:
        return self.name

    # Provide a regular expression string based on detection indicators.
    def fmtRegex(self):
        self.regStr = ".*(" + "|".join(self.indicators) + ")$"
    
    # Sets rule indicators based on INDICATOR_PATH.
    def setFilterRegex(self) -> None:
        # Open indicators
        try:
            lst = open(INDICATOR_PATH)
        except Exception as e:
             logging.error("failed to open file %s: %s", INDICATOR_PATH, str(e))
             raise e
        
        for indicator in lst:
            self.indicators.append(indicator.strip().split(",")[0])

        lst.close()

    # Run detection logic.
    def run(self) -> list[spb.Signal]:
        # Set regex for filtering.
        try:
            self.setFilterRegex()
        except Exception as e:
            raise e
        
        # <TODO: Implement me!>
        # Find any logs that contain indicators of compromise from indicatorPath:
        #   1. Filter logs to what is relevant, then
        #   2. [Optional] Aggregate logs based on source IP address.
        #   3. Return the set of interesting logs as a list of spb.Signal 
        # 
        # Expected output:
        # Option #1: If the aggregation step is skipped, the list of spb.Signal returned should have `bad_domain_filtered` set.
        # Option #2: If both filtering and aggregation are performed, the list of spb.Signal returned should have `bad_domain` set.
        # 
        # Hint #1: import re, and use self.regStr.
        # Hint #2: Aggregation is easier using a dict!
        # Hint #3: Check the fields you need to populate by inspecting the spb.BadDomain protobuf message.
        return None