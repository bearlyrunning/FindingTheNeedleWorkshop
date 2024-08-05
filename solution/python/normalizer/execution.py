import csv
import logging
import normalizer
import validator
from generated import normalized_log_pb2 as nlpb

class ExecutionNormalizer(normalizer.NormalizerInterface):

    # Constructor.
    def __init__(self, input="", binaryOutput="", jsonOutput=""):
        self.input = input
        self.binaryOutput = binaryOutput
        self.jsonOutput=jsonOutput

    # Return input file location.
    def getInput(self) -> str:
        return self.input
    
    # Return binary file location.
    def getBinaryOutput(self) -> str:
        return self.binaryOutput
    
    # Return JSON file location.
    def getJsonOutput(self) -> str:
        return self.jsonOutput
    
    # Normalize log source line-by-line.
    def normalize(self, line="") -> nlpb.NormalizedLog:
        # Initialise protobuf log structures
        log = nlpb.NormalizedLog()

        # Parse line using csv library due to presence of commas in data
        fields = list(csv.reader([line]))[0]

        if len(fields) != 9:
            logging.warning("invalid number of fields found; expect 9, found %d: %s",len(fields), line)
            return None
        
        # Parse and populate timestamp
        try:
            ts = validator.validateTimestamp(fields[0])
            log.execution_log.timestamp.CopyFrom(ts)
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        # Parse and populate filepath
        log.execution_log.filepath = fields[1].strip("\"")

        # Parse and populate command
        log.execution_log.command = fields[2].strip("\"")

        # Parse and populate uid
        try:
            uid = validator.validateInt(fields[3])
            log.execution_log.uid = uid
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # Parse and populate pid
        try:
            pid = validator.validateInt(fields[4])
            log.execution_log.pid = pid
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        # Parse and populate ppid
        try:
            ppid = validator.validateInt(fields[5])
            log.execution_log.ppid = ppid
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        # Parse and populate cwd
        log.execution_log.cwd = fields[6].strip("\"")

        # Parse and populate hostname
        log.execution_log.hostname = fields[7].strip("\"")

        # Parse and populate platform
        try:
            log.execution_log.platform = validator.validatePlatform(fields[8])
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        return log