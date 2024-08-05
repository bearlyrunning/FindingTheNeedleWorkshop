To build and run the normalizer binary:

```
$ go build
$ ./normalizer
```

To run the unit tests:

```
go test -v -run TestDNSNormalizer github.com/bearlyrunning/FindingTheNeedle/go/normalizer
go test -v -run TestNetflowNormalizer github.com/bearlyrunning/FindingTheNeedle/go/normalizer
go test -v -run TestExecutionNormalizer github.com/bearlyrunning/FindingTheNeedle/go/normalizer
go test -v -run TestSplitWithEscape github.com/bearlyrunning/FindingTheNeedle/go/normalizer
```
