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
    
    # normalize() function does the following:
    # - reads each `line` of /data/dns/dns.log
    # - parses and validates each comma separated field in the log
    # - output the log as a NormalizedLog proto message which eventually get saved as:
    #   - /data/dns/dns.binpb
    #   - /data/dns/dns.json
    def normalize(self, line="") -> nlpb.NormalizedLog:
        # Initialise protobuf log structures
        log = nlpb.NormalizedLog()

        # Validate number of fields in each log line.
        fields = line.split(",")
        if len(fields) != 8:
            logging.warning("invalid number of fields found; expect 8, found %d: %s", len(fields), line)
            return None
        
        # <TODO: Implement me!>
        # Implement the missing validate function(s) in validator.py file.
        # Use validator functions below (e.g. validator.validateTime(...))
        try:
            # <TODO: Implement me!>
            # Parse and populate timestamp using validateTime() in validator.py.
            ts = validator.validateTime("TODO: Implement me!")
            log.dns_log.timestamp.CopyFrom(ts)
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # <TODO: Implement me!>
        # Set the log_source field.
        log.dns_log.log_source = "TODO: Implement me!"
        
        try:
            # <TODO: Implement me!>
            # Parse and populate src IP using validateIP() in validator.py.
            src_ip = validator.validateIP("TODO: Implement me!")
            log.dns_log.source_ip = src_ip
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate resolver IP using validateIP() in validator.py.
            resolver_ip = validator.validateIP("TODO: Implement me!")
            log.dns_log.resolver_ip = resolver_ip
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # TODO: <Implement me!>
            # Parse and populate query using validateQuery() in validator.py.
            log.dns_log.query = validator.validateQuery("TODO: Implement me!")
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # TODO: <Implement me!>
        # Set DNS query type
        log.dns_log.type = "TODO: Implement me!"

        # TODO: <Implement me!>
        # Set DNS query answer
        log.dns_log.answer = "TODO: Implement me!"

        try:
            # TODO: <Implement me!>
            # Parse and populate return code using validateReturnCode() in validator.py.
            log.dns_log.return_code = validator.validateReturnCode("TODO: Implement me!")
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        return log