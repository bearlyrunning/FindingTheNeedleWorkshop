import logging
import normalizer
import validator
from generated import normalized_log_pb2 as nlpb

class DNSNormalizer(normalizer.NormalizerInterface):
    
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
        fields = line.split(",")
        if len(fields) != 8:
            logging.warning("invalid number of fields found; expect 8, found %d: %s", len(fields), line)
            return None
        
        # Parse and populate timestamp
        try:
            ts = validator.validateTime(fields[0])
            log.dns_log.timestamp.CopyFrom(ts)
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # Set log source
        log.dns_log.log_source = fields[1]
        
        # Parse and populate src IP
        try:
            src_ip = validator.validateIP(fields[2])
            log.dns_log.source_ip = src_ip
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # Parse and populate resolver IP
        try:
            resolver_ip = validator.validateIP(fields[3])
            log.dns_log.resolver_ip = resolver_ip
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # Parse and populate query
        try:
            log.dns_log.query = validator.validateQuery(fields[4])
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # Set DNS query type
        log.dns_log.type = fields[5]

        # Set DNS query answer
        log.dns_log.answer = fields[6]

        # Parse and populate return code (fields[7])
        try:
            log.dns_log.return_code = validator.validateReturnCode(fields[7])
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        return log