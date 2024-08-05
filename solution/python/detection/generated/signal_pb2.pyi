from google.protobuf import timestamp_pb2 as _timestamp_pb2
from . import normalized_log_pb2 as _normalized_log_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Signal(_message.Message):
    __slots__ = ("bad_domain_filtered", "bad_domain", "dns_tunnel", "browser_sub_proc")
    BAD_DOMAIN_FILTERED_FIELD_NUMBER: _ClassVar[int]
    BAD_DOMAIN_FIELD_NUMBER: _ClassVar[int]
    DNS_TUNNEL_FIELD_NUMBER: _ClassVar[int]
    BROWSER_SUB_PROC_FIELD_NUMBER: _ClassVar[int]
    bad_domain_filtered: _normalized_log_pb2.DNS
    bad_domain: BadDomain
    dns_tunnel: DNSTunnel
    browser_sub_proc: BrowserSubProc
    def __init__(self, bad_domain_filtered: _Optional[_Union[_normalized_log_pb2.DNS, _Mapping]] = ..., bad_domain: _Optional[_Union[BadDomain, _Mapping]] = ..., dns_tunnel: _Optional[_Union[DNSTunnel, _Mapping]] = ..., browser_sub_proc: _Optional[_Union[BrowserSubProc, _Mapping]] = ...) -> None: ...

class BadDomain(_message.Message):
    __slots__ = ("timestamp_start", "timestamp_end", "bad_domain", "source_ip", "hostname", "dns_log")
    TIMESTAMP_START_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_END_FIELD_NUMBER: _ClassVar[int]
    BAD_DOMAIN_FIELD_NUMBER: _ClassVar[int]
    SOURCE_IP_FIELD_NUMBER: _ClassVar[int]
    HOSTNAME_FIELD_NUMBER: _ClassVar[int]
    DNS_LOG_FIELD_NUMBER: _ClassVar[int]
    timestamp_start: _timestamp_pb2.Timestamp
    timestamp_end: _timestamp_pb2.Timestamp
    bad_domain: str
    source_ip: str
    hostname: str
    dns_log: _containers.RepeatedCompositeFieldContainer[_normalized_log_pb2.DNS]
    def __init__(self, timestamp_start: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., timestamp_end: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., bad_domain: _Optional[str] = ..., source_ip: _Optional[str] = ..., hostname: _Optional[str] = ..., dns_log: _Optional[_Iterable[_Union[_normalized_log_pb2.DNS, _Mapping]]] = ...) -> None: ...

class DNSTunnel(_message.Message):
    __slots__ = ("timestamp_start", "timestamp_end", "tunnel_ip", "source_ip", "hostname", "bytes_in_total", "bytes_out_total", "netflow_log")
    TIMESTAMP_START_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_END_FIELD_NUMBER: _ClassVar[int]
    TUNNEL_IP_FIELD_NUMBER: _ClassVar[int]
    SOURCE_IP_FIELD_NUMBER: _ClassVar[int]
    HOSTNAME_FIELD_NUMBER: _ClassVar[int]
    BYTES_IN_TOTAL_FIELD_NUMBER: _ClassVar[int]
    BYTES_OUT_TOTAL_FIELD_NUMBER: _ClassVar[int]
    NETFLOW_LOG_FIELD_NUMBER: _ClassVar[int]
    timestamp_start: _timestamp_pb2.Timestamp
    timestamp_end: _timestamp_pb2.Timestamp
    tunnel_ip: str
    source_ip: str
    hostname: str
    bytes_in_total: int
    bytes_out_total: int
    netflow_log: _containers.RepeatedCompositeFieldContainer[_normalized_log_pb2.Netflow]
    def __init__(self, timestamp_start: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., timestamp_end: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., tunnel_ip: _Optional[str] = ..., source_ip: _Optional[str] = ..., hostname: _Optional[str] = ..., bytes_in_total: _Optional[int] = ..., bytes_out_total: _Optional[int] = ..., netflow_log: _Optional[_Iterable[_Union[_normalized_log_pb2.Netflow, _Mapping]]] = ...) -> None: ...

class BrowserSubProc(_message.Message):
    __slots__ = ("execution", "source_ip")
    EXECUTION_FIELD_NUMBER: _ClassVar[int]
    SOURCE_IP_FIELD_NUMBER: _ClassVar[int]
    execution: _normalized_log_pb2.Execution
    source_ip: str
    def __init__(self, execution: _Optional[_Union[_normalized_log_pb2.Execution, _Mapping]] = ..., source_ip: _Optional[str] = ...) -> None: ...
