## Dependencies
* https://cloud.google.com/sdk/docs/install-sdk
* https://grpc.io/docs/languages/go/quickstart/

### Compile service definition

#### Golang

```
$ protoc \
--proto_path=./service/ \
--go_out=./service/ --go_opt=paths=source_relative \
--go-grpc_out=./service/ --go-grpc_opt=paths=source_relative \
enrichment.proto
```

```
$ protoc \
--proto_path=./service/ \
--go_out=../go/generated/enrichmentpb --go_opt=paths=source_relative \
--go-grpc_out=../go/generated/enrichmentpb --go-grpc_opt=paths=source_relative \
enrichment.proto
```

#### Python

```
python3 -m grpc_tools.protoc \
-Iservice/ \
--python_out=../python/enrichment/generated/ \
--pyi_out=../python/enrichment/generated/ \
--grpc_python_out=../python/enrichment/generated/ \
enrichment.proto
```

**Note:** The generated Python gRPC module (`enrichment_pb2_grpc.py`) needs to be manually amended to fix a broken import. 

The following line needs to be amended:

```
import enrichment_pb2 as enrichment__pb2
```

To the following:

```
from . import enrichment_pb2 as enrichment__pb2
```

Further information is available in [this GitHub-reported issue](https://github.com/protocolbuffers/protobuf/issues/1491).

### Running the server locally
```
$ cd FindingTheNeedle/grpc/server
$ go build
$ ./server
```
