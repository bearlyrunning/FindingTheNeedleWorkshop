## Doc
https://protobuf.dev/programming-guides/proto3/#generating

## Golang

```
$ protoc \
--proto_path=./ \
--go_out=../go/generated/normalizedlogpb --go_opt=paths=source_relative \
normalized_log.proto

$ protoc \
--proto_path=./ \
--go_out=../go/generated/signalpb --go_opt=paths=source_relative \
signal.proto
```

## Python

```
$ protoc -I=./ \
  --python_out=../python/normalizer/generated/ \
  --pyi_out=../python/normalizer/generated \
  normalized_log.proto

$ protoc -I=./ \
  --python_out=../python/detection/generated/ \
  --pyi_out=../python/detection/generated \
  normalized_log.proto

$ protoc -I=./ \
  --python_out=../python/detection/generated/ \
  --pyi_out=../python/detection/generated \
  signal.proto

  $ protoc -I=./ \
  --python_out=../python/enrichment/generated/ \
  --pyi_out=../python/enrichment/generated \
  signal.proto

$ protoc -I=./ \
  --python_out=../python/enrichment/generated/ \
  --pyi_out=../python/enrichment/generated \
  signal.proto
```