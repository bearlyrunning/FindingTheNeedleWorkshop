import abc
from generated import signal_pb2 as spb

"""Detection interface for class-specific detection logic."""
class DetectionInterface(metaclass=abc.ABCMeta):
    @classmethod
    def __subclasshook__(cls, subclass):
        return (hasattr(subclass, 'Run') and 
                callable(subclass.Run)  or 
                NotImplemented)
    
    @abc.abstractmethod
    def run(self) -> list[spb.Signal]:
        """Run implemented detection logic."""
        raise NotImplementedError
    
    @abc.abstractmethod
    def ruleName(self) -> str:
        """Return detection rule name."""
        raise NotImplementedError