from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class NormalizedLog(_message.Message):
    __slots__ = ("dns_log", "netflow_log", "execution_log")
    DNS_LOG_FIELD_NUMBER: _ClassVar[int]
    NETFLOW_LOG_FIELD_NUMBER: _ClassVar[int]
    EXECUTION_LOG_FIELD_NUMBER: _ClassVar[int]
    dns_log: DNS
    netflow_log: Netflow
    execution_log: Execution
    def __init__(self, dns_log: _Optional[_Union[DNS, _Mapping]] = ..., netflow_log: _Optional[_Union[Netflow, _Mapping]] = ..., execution_log: _Optional[_Union[Execution, _Mapping]] = ...) -> None: ...

class DNS(_message.Message):
    __slots__ = ("timestamp", "query", "type", "answer", "return_code", "source_ip", "resolver_ip", "log_source")
    class ReturnCode(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        UNSPECIFIED: _ClassVar[DNS.ReturnCode]
        NOERROR: _ClassVar[DNS.ReturnCode]
        FORMERR: _ClassVar[DNS.ReturnCode]
        SERVFAIL: _ClassVar[DNS.ReturnCode]
        NXDOMAIN: _ClassVar[DNS.ReturnCode]
        NOTIMP: _ClassVar[DNS.ReturnCode]
        REFUSED: _ClassVar[DNS.ReturnCode]
        YXDOMAIN: _ClassVar[DNS.ReturnCode]
        XRRSET: _ClassVar[DNS.ReturnCode]
        NOTAUTH: _ClassVar[DNS.ReturnCode]
        NOTZONE: _ClassVar[DNS.ReturnCode]
    UNSPECIFIED: DNS.ReturnCode
    NOERROR: DNS.ReturnCode
    FORMERR: DNS.ReturnCode
    SERVFAIL: DNS.ReturnCode
    NXDOMAIN: DNS.ReturnCode
    NOTIMP: DNS.ReturnCode
    REFUSED: DNS.ReturnCode
    YXDOMAIN: DNS.ReturnCode
    XRRSET: DNS.ReturnCode
    NOTAUTH: DNS.ReturnCode
    NOTZONE: DNS.ReturnCode
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    QUERY_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    ANSWER_FIELD_NUMBER: _ClassVar[int]
    RETURN_CODE_FIELD_NUMBER: _ClassVar[int]
    SOURCE_IP_FIELD_NUMBER: _ClassVar[int]
    RESOLVER_IP_FIELD_NUMBER: _ClassVar[int]
    LOG_SOURCE_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    query: str
    type: str
    answer: str
    return_code: DNS.ReturnCode
    source_ip: str
    resolver_ip: str
    log_source: str
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., query: _Optional[str] = ..., type: _Optional[str] = ..., answer: _Optional[str] = ..., return_code: _Optional[_Union[DNS.ReturnCode, str]] = ..., source_ip: _Optional[str] = ..., resolver_ip: _Optional[str] = ..., log_source: _Optional[str] = ...) -> None: ...

class Netflow(_message.Message):
    __slots__ = ("timestamp", "src_ip", "src_port", "dst_ip", "dst_port", "bytes_in", "bytes_out", "packets_in", "packets_out", "protocol", "log_source")
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    SRC_IP_FIELD_NUMBER: _ClassVar[int]
    SRC_PORT_FIELD_NUMBER: _ClassVar[int]
    DST_IP_FIELD_NUMBER: _ClassVar[int]
    DST_PORT_FIELD_NUMBER: _ClassVar[int]
    BYTES_IN_FIELD_NUMBER: _ClassVar[int]
    BYTES_OUT_FIELD_NUMBER: _ClassVar[int]
    PACKETS_IN_FIELD_NUMBER: _ClassVar[int]
    PACKETS_OUT_FIELD_NUMBER: _ClassVar[int]
    PROTOCOL_FIELD_NUMBER: _ClassVar[int]
    LOG_SOURCE_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    src_ip: str
    src_port: int
    dst_ip: str
    dst_port: int
    bytes_in: int
    bytes_out: int
    packets_in: int
    packets_out: int
    protocol: str
    log_source: str
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., src_ip: _Optional[str] = ..., src_port: _Optional[int] = ..., dst_ip: _Optional[str] = ..., dst_port: _Optional[int] = ..., bytes_in: _Optional[int] = ..., bytes_out: _Optional[int] = ..., packets_in: _Optional[int] = ..., packets_out: _Optional[int] = ..., protocol: _Optional[str] = ..., log_source: _Optional[str] = ...) -> None: ...

class Execution(_message.Message):
    __slots__ = ("timestamp", "filepath", "command", "uid", "pid", "ppid", "cwd", "hostname", "platform")
    class Platform(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        UNSPECIFIED: _ClassVar[Execution.Platform]
        LINUX: _ClassVar[Execution.Platform]
        WINDOWS: _ClassVar[Execution.Platform]
        MAC: _ClassVar[Execution.Platform]
    UNSPECIFIED: Execution.Platform
    LINUX: Execution.Platform
    WINDOWS: Execution.Platform
    MAC: Execution.Platform
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    FILEPATH_FIELD_NUMBER: _ClassVar[int]
    COMMAND_FIELD_NUMBER: _ClassVar[int]
    UID_FIELD_NUMBER: _ClassVar[int]
    PID_FIELD_NUMBER: _ClassVar[int]
    PPID_FIELD_NUMBER: _ClassVar[int]
    CWD_FIELD_NUMBER: _ClassVar[int]
    HOSTNAME_FIELD_NUMBER: _ClassVar[int]
    PLATFORM_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    filepath: str
    command: str
    uid: int
    pid: int
    ppid: int
    cwd: str
    hostname: str
    platform: Execution.Platform
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., filepath: _Optional[str] = ..., command: _Optional[str] = ..., uid: _Optional[int] = ..., pid: _Optional[int] = ..., ppid: _Optional[int] = ..., cwd: _Optional[str] = ..., hostname: _Optional[str] = ..., platform: _Optional[_Union[Execution.Platform, str]] = ...) -> None: ...
