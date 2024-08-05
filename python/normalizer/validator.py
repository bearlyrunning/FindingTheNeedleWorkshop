from google.protobuf.timestamp_pb2 import Timestamp

def validateTime(time) -> Timestamp:
    # <TODO: Implement me!>
    # Parse a datetime string (format: "%Y-%m-%d %H:%M:%S.%f") to a Timestamp proto message.
    # Raise an Exception if the datetime string is not valid.
    # Fix the placeholder return below.
    # Hint #1: import datetime
    # Hint #2: make use of the google.protobuf.timestamp_pb2 module.
    return None

def validateTimestamp(time) -> Timestamp:
    # <TODO: Implement me!>
    # Parse a epoch timestamp string to a Timestamp proto message.
    # Raise an Exception if the timestamp string is not valid.
    # Fix the placeholder return below.
    # Hint #1: what field(s) does the Timestamp proto message contain?
    return None

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
    # <TODO: Implement me!>
    # Convert return code to proto ENUM nlpb.DNS.ReturnCode.
    # Confirm the string is a valid return code (hint: check the range of the enum).
    # Raise an Exception if the code is not valid.
    # Don't forget to increment return code by 1 as the enum value 0 is reserved for default value (e.g. unspecified) only.
    # Hint: Hint: check the auto-generated normalized_log_pb2 package for suitable conversion approach.
    return None

def validateInt(s) -> int:
    # <TODO: Implement me!>
    # Convert the number, provided as a string, to an integer.
    # Raise an Exception if the string cannot be converted to a valid integer.
    return None

def validatePlatform(s) -> str:
    # <TODO: Implement me!>
    # Convert platform string to Platform ENUM.
    # If the platform string is not valid, raise an Exception.
    # Hint: check the auto-generated normalized_log_pb2 package for suitable conversion approach.
    return None