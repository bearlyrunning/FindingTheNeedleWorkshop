from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class IP(_message.Message):
    __slots__ = ("ip",)
    IP_FIELD_NUMBER: _ClassVar[int]
    ip: str
    def __init__(self, ip: _Optional[str] = ...) -> None: ...

class Host(_message.Message):
    __slots__ = ("name", "platform")
    class Platform(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        UNSPECIFIED: _ClassVar[Host.Platform]
        LINUX: _ClassVar[Host.Platform]
        WINDOWS: _ClassVar[Host.Platform]
        MAC: _ClassVar[Host.Platform]
    UNSPECIFIED: Host.Platform
    LINUX: Host.Platform
    WINDOWS: Host.Platform
    MAC: Host.Platform
    NAME_FIELD_NUMBER: _ClassVar[int]
    PLATFORM_FIELD_NUMBER: _ClassVar[int]
    name: str
    platform: Host.Platform
    def __init__(self, name: _Optional[str] = ..., platform: _Optional[_Union[Host.Platform, str]] = ...) -> None: ...
