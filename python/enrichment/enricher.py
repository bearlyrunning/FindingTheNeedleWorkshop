import abc
from generated import signal_pb2 as spb

"""Enrichment interface for detection enrichments."""
class EnricherInterface(metaclass=abc.ABCMeta):
    @classmethod
    def __subclasshook__(cls, subclass):
        return (hasattr(subclass, 'load') and 
                callable(subclass.load) and 
                hasattr(subclass, 'enrich') and 
                callable(subclass.enrich) and 
                hasattr(subclass, 'output') and 
                callable(subclass.output) or
                NotImplemented)
    
    @abc.abstractmethod
    def getName(self) -> str:
        """Return enrichment name."""
        raise NotImplementedError
    
    @abc.abstractmethod
    def getSignals(self) -> list[spb.Signal]:
        """Return signals associated with an Enricher."""
        raise NotImplementedError
    
    @abc.abstractmethod
    def setSignals(self, signals) -> None:
        """Assign signals associated with an Enricher."""
        raise NotImplementedError