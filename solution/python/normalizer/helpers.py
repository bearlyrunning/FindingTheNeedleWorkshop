# Check single protobuf equivalence via string serialization.
def checkProtoEqual(want, got) -> bool:
    return want.SerializeToString(deterministic=True) == got.SerializeToString(deterministic=True)
    
# Check protobuf list equivalence via string serialization.
def checkProtoListEqual(want, got)  -> bool:
    want_serialized = []
    got_serialised = []
    for i in want:
        want_serialized.append(i.SerializeToString(deterministic=True))
    for i in got:
        got_serialised.append(i.SerializeToString(deterministic=True))
    return want_serialized == got_serialised   