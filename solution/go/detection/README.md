To build and run the detection binary:

```
$ go build
$ ./detection
```

To run the unit tests:

```
go test -v -run TestFmtRegex github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestBadDomainFilter github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestBadDomainAggregate github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestDNSTunnelFilter github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestDNSTunnelAggregate github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestBrowserSubProcDetection github.com/bearlyrunning/FindingTheNeedle/go/detection
```
