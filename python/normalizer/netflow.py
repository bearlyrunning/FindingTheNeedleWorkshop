import logging
import normalizer
import validator
from generated import normalized_log_pb2 as nlpb

class NetflowNormalizer(normalizer.NormalizerInterface):

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

        if len(fields) != 11:
            logging.warning("invalid number of fields found; expect 11, found %d: %s", len(fields), line)
            return None
        
        try:
            # <TODO: Implement me!>
            # Parse and populate timestamp using validateTime() validator.py.
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        # <TODO: Implement me!>
        # Parse and populate log source
        # Parse and populate protocol

        try:
            # <TODO: Implement me!>
            # Parse and populate source IP address using validateIP() in validator.py.
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate src_port using validatePort() in validator.py
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        try:
            # <TODO: Implement me!>
            # Parse and populate dst_ip using validateIP() in validator.py
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        try:
            # <TODO: Implement me!>
            # Parse and populate dst_port using validatePort() in validator.py
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        try:
            # <TODO: Implement me!>
            # Parse and populate bytes_in using validateInt() in validator.py
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate bytes_out using validateInt() in validator.py
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate packets_in using validateInt() in validator.py
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate packets_out using validateInt() in validator.py
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        return log