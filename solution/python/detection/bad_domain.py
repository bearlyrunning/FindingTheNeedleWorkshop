import detection
import logging
import re
from generated import normalized_log_pb2 as nlpb
from generated import signal_pb2 as spb
import google.protobuf.timestamp_pb2 as tspb

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

    # Match indicator presence in logs.
    def filter(self) -> list[nlpb.DNS]:
        self.setFilterRegex
        self.fmtRegex()
        matched = []
        
        for log in self.logs.dns:
            if re.search(self.regStr, log.query) or re.search(self.regStr, log.answer):
                matched.append(log)

        return matched
    
    # Comput signal aggregation based on source IP address.
    def aggregate(self, matched=[]) -> list[spb.Signal]:
        sigs = []
        # Initialise a dict of source_ip -> list[nlpb.DNS]
        keyBySrc = {}

        # Aggregate by source_ip.
        for m in matched:
            if m.source_ip not in keyBySrc:
                keyBySrc[m.source_ip] = [m]
            else:
                keyBySrc[m.source_ip].append(m)

        for src, logs in keyBySrc.items():
            # Determine time window of aggregated logs.
            earliest = logs[0].timestamp.ToDatetime()
            latest = logs[0].timestamp.ToDatetime()
            for l in logs[1:]:
                if l.timestamp.ToDatetime() < earliest:
                    earliest = l.timestamp.ToDatetime()
                    continue
                if l.timestsamp.ToDateTime() > latest:
                    latest = l.timestamp.ToDatetime()
            ts_earliest = tspb.Timestamp()
            ts_earliest.FromDatetime(earliest)
            ts_latest = tspb.Timestamp()
            ts_latest.FromDatetime(latest)

            # Find the bad indicator quried
            m = re.findall(self.regStr, logs[0].query)
            if len(m) == 0:
                m = re.findall(self.regStr, logs[0].answer)

            # Create aggregated entry per source_ip.
            sigs.append(spb.Signal(
                bad_domain=spb.BadDomain(
                    timestamp_start=ts_earliest,
                    timestamp_end=ts_latest,
                    source_ip=src,
                    bad_domain=m[0],
                    dns_log=logs
                )
            ))

        return sigs

    # Run detection logic.
    def run(self) -> list[spb.Signal]:
        # Set regex for filtering.
        try:
            self.setFilterRegex()
        except Exception as e:
            raise e
        
        # Find any logs that contain indicators of compromise.
        # Run detection logic: 1) filter 2) aggregate.
        return self.aggregate(self.filter())