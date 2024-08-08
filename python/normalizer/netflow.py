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
    
    # normalize() function does the following:
    # - reads each `line` of /data/netflow/netflow.log
    # - parses and validates each comma separated field in the log
    # - output the log as a NormalizedLog proto (see /proto/normalized.proto) message which eventually get saved as:
    #   - /data/netflow/netflow_normalized.binpb
    #   - /data/netflow/netflow_normalized.json
    def normalize(self, line="") -> nlpb.NormalizedLog:
        # Initialise protobuf log structures
        log = nlpb.NormalizedLog()
        fields = line.split(",")

        if len(fields) != 11:
            logging.warning("invalid number of fields found; expect 11, found %d: %s", len(fields), line)
            return None
        
        # <TODO: Implement me!>
        # Implement the missing validate function(s) in validator.py file.
        try:
            # <TODO: Implement me!>
            # Parse and populate timestamp using validateTime() validator.py.
            ts = validator.validateTime("TODO: Implement me!")
            log.netflow_log.timestamp.CopyFrom(ts)
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        # <TODO: Implement me!>
        # Parse and populate log source
        log.netflow_log.log_source = "TODO: Implement me!"

        # <TODO: Implement me!>
        # Parse and populate protocol
        log.netflow_log.protocol = "TODO: Implement me!"

        try:
            # <TODO: Implement me!>
            # Parse and populate source IP address using validateIP() in validator.py.
            src_ip = validator.validateIP("TODO: Implement me!")
            log.netflow_log.src_ip = src_ip
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate src_port using validatePort() in validator.py
            src_port = validator.validatePort("TODO: Implement me!")
            log.netflow_log.src_port = src_port
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        try:
            # <TODO: Implement me!>
            # Parse and populate dst_ip using validateIP() in validator.py
            dst_ip = validator.validateIP("TODO: Implement me!")
            log.netflow_log.dst_ip = dst_ip
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        try:
            # <TODO: Implement me!>
            # Parse and populate dst_port using validatePort() in validator.py
            dst_port = validator.validatePort("TODO: Implement me!")
            log.netflow_log.dst_port = dst_port
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        try:
            # <TODO: Implement me!>
            # Parse and populate bytes_in using validateInt() in validator.py
            bytes_in = validator.validateInt("TODO: Implement me!")
            log.netflow_log.bytes_in = bytes_in
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate bytes_out using validateInt() in validator.py
            bytes_out = validator.validateInt("TODO: Implement me!")
            log.netflow_log.bytes_out = bytes_out
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate packets_in using validateInt() in validator.py
            packets_in = validator.validateInt("TODO: Implement me!")
            log.netflow_log.packets_in = packets_in
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        try:
            # <TODO: Implement me!>
            # Parse and populate packets_out using validateInt() in validator.py
            packets_out = validator.validateInt("TODO: Implement me!")
            log.netflow_log.packets_out = packets_out
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        return log