import abc
from generated import normalized_log_pb2 as nlpb

"""Normalizer interface for DNS, Netflow, and Execution normalizers."""
class NormalizerInterface(metaclass=abc.ABCMeta):
    @classmethod
    def __subclasshook__(cls, subclass):
        return (hasattr(subclass, 'getInput') and 
                callable(subclass.getInput) and 
                hasattr(subclass, 'getBinaryOutput') and 
                callable(subclass.getBinaryOutput) and
                hasattr(subclass, 'getJsonOutput') and 
                callable(subclass.getJsonOutput) and
                hasattr(subclass, 'normalize') and 
                callable(subclass.normalize) or 
                NotImplemented)
    
    @abc.abstractmethod
    def getInput(self) -> str:
        """Return input file location."""
        raise NotImplementedError
    
    @abc.abstractmethod
    def getBinaryOutput(self) -> str:
        """Return binary file location."""
        raise NotImplementedError
    
    @abc.abstractmethod
    def getJsonOutput(self) -> str:
        """Return JSON file location."""
        raise NotImplementedError

    @abc.abstractmethod
    def normalize(self, line) -> nlpb.NormalizedLog:
        """Parse all log lines and output formatted data."""
        raise NotImplementedError
