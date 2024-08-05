from datetime import datetime
from generated import normalized_log_pb2 as nlpb
from google.protobuf.timestamp_pb2 import Timestamp
import ipaddress

def validateTime(time) -> Timestamp:
    try:
        t = datetime.strptime(time, "%Y-%m-%d %H:%M:%S.%f")
    except ValueError as e:
        raise e
    ts = Timestamp()
    ts.FromDatetime(t)
    return ts

def validateTimestamp(time) -> Timestamp:
    try:
        t = int(time)
    except Exception as e:
        raise e
    ts = Timestamp()
    ts.FromSeconds(t)
    return ts

def validateIP(ip) -> str:
    try:
        addr = ipaddress.ip_address(ip)
    except ValueError as e:
        raise e
    return str(addr)

def validatePort(s) -> int:
    try:
        port = int(s)
    except Exception as e:
        raise e
    
    if port < 0 or port > 65535:
        raise ValueError("unexpected port number found: %s", s)
    
    return port

def validateQuery(s) -> str:
    if s == "":
        raise ValueError("empty query found")
    return s

def validateReturnCode(s) -> str:
    try:
        i = int(s)
    except Exception as e:
        raise e
    code = nlpb.DNS.ReturnCode.Name(i + 1)
    return code

def validateInt(s) -> int:
    try:
        i = int(s)
    except Exception as e:
        raise e
    return i

def validatePlatform(s) -> str:
    try: 
        p = nlpb.Execution.Platform.Value(s.strip("\n").strip("\""))
    except Exception as e:
        raise e
    return nlpb.Execution.Platform.Name(p)