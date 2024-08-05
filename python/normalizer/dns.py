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
        
        # <TODO> Implement me!>
        # Implement the validate functions in validator.py.
        # Use validator functions below (e.g. validator.validateTime(...))
        try:
            # <TODO: Implement me!>
            # Parse and populate timestamp using validateTime() in validator.py.
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # <TODO: Implement me!>
        # Set the log_source field.
        
        try:
            # <TODO: Implement me!>
            # Parse and populate src IP using validateIP() in validator.py.
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate resolver IP using validateIP() in validator.py.
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # TODO: <Implement me!>
            # Parse and populate query using validateQuery() in validator.py.
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # TODO: <Implement me!>
        # Set DNS query type
        # Set DNS query answer

        try:
            # TODO: <Implement me!>
            # Parse and populate return code using validateReturnCode() in validator.py.
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        return log