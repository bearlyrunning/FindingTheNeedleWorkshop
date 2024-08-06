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
    # <TODO: Implement me!>
    # Confirm the IP string contains a valid IP address.
    # Raise an Exception if the IP is not valid.
    # Return back the valid IP as a string.
    return None

def validatePort(s) -> int:
    # <TODO: Implement me!>
    # Convert the port, passed as a string, to an integer.
    # Confirm the string is a valid port number.
    # Raise an Exception if the port cannot be converted to a string.
    # Raise a ValueError("unexpected port number found: %s", s) if the port is not valid.
    # Return the port as an integer.
    return None

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
    # <TODO: Implement me!>
    # Convert the number, provided as a string, to an integer.
    # Raise an Exception if the string cannot be converted to a valid integer.
    return None

def validatePlatform(s) -> str:
    try: 
        p = nlpb.Execution.Platform.Value(s.strip("\n").strip("\""))
    except Exception as e:
        raise e
    return nlpb.Execution.Platform.Name(p)