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
        
        # <TODO: Implement me!>
        # Parse and populate timestamp using validateTimestamp() in validator.py.

        # <TODO: Implement me!>
        # Parse and populate filepath
        # Parse and populate command

        # <TODO: Implement me!>
        # Parse and populate uid using validateInt() in validator.py.

        # <TODO: Implement me!>
        # Parse and populate pid using validateInt() in validator.py.

        # <TODO: Implement me!>
        # Parse and populate ppid using validateInt() in validator.py.

        
        # <TODO: Implement me!>
        # Parse and populate cwd
        # Parse and populate hostname

        # <TODO: Implement me!>
        # Parse and populate platform using validatePlatform() in validator.py.
   
        return log