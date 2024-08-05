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
        
        # Parse and populate timestamp
        try:
            ts = validator.validateTime(fields[0])
            log.netflow_log.timestamp.CopyFrom(ts)
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        # Parse and populate log source
        log.netflow_log.log_source = fields[1]

        # Parse and populate protocol
        log.netflow_log.protocol = fields[2]

        # Parse and populate source IP address
        try:
            src_ip = validator.validateIP(fields[3])
            log.netflow_log.src_ip = src_ip
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # Parse and populate src_port
        try:
            src_port = validator.validatePort(fields[4])
            log.netflow_log.src_port = src_port
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        # Parse and populate dst_ip
        try:
            dst_ip = validator.validateIP(fields[5])
            log.netflow_log.dst_ip = dst_ip
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        # Parse and populate dst_port
        try:
            dst_port = validator.validatePort(fields[6])
            log.netflow_log.dst_port = dst_port
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None
        
        # Parse and populate bytes_in
        try:
            bytes_in = validator.validateInt(fields[7])
            log.netflow_log.bytes_in = bytes_in
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # Parse and populate bytes_out
        try:
            bytes_out = validator.validateInt(fields[8])
            log.netflow_log.bytes_out = bytes_out
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # Parse and populate packets_in
        try:
            packets_in = validator.validateInt(fields[9])
            log.netflow_log.packets_in = packets_in
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        # Parse and populate packets_out
        try:
            packets_out = validator.validateInt(fields[10])
            log.netflow_log.packets_out = packets_out
        except Exception as err:
            logging.warning("%s, skipping: %s", err, line)
            return None

        return log