# Detection Logic in Golang

In this section, your task is, for each of the bad DNS resolution, DNS tunneling, and browser subprocess detections:

1. Design the detection logic for each detection rule, and implement the `run()` function.
2. Implement the unimplemented unit tests.
3. Make sure your unit tests run as expected.

The locations you need to make code changes can be found by looking for the following comment string:

```
<TODO: Implement me!>
```

These changes need to take place in separate files. We recommend completing this section in the following order:

1. `bad_domain.go`, with unit tests in `bad_domain_test.go`.
2. `dns_tunnel.go`, with unit tests in `dns_tunnel_test.go`.
3. `browser_sub_proc.go`, with unit tests in `browser_sub_proc_test.go`.

## Input and Output locations

The following table outlines where each detection rule ingests data from, and where its output is written to:

| **Rule**                         | **Code Location**     | **Input Data Location**                           | **Output Data Location**                  |
|----------------------------------|-----------------------|---------------------------------------------------|-------------------------------------------|
| Bad DNS domain resolution        | `bad_domain.go`       | `../../data/dns/dns_normalized.binpb"`            | `../../data/signal/bad_domain.json`       |
| DNS tunnelling                   | `dns_tunnel.go`       | `../../data/netflow/netflow_normalized.binpb`     | `../../data/signal/dns_tunnel.json`       |
| Suspicious browser child process | `browser_sub_proc.go` | `../../data/execution/execution_normalized.binpb` | `../../data/signal/browser_sub_proc.json` |

## Running your detection logic

To build and run the detection binary:

```
$ go build
$ ./detection
```

## Running the unit tests

To run the unit tests:

```
go test -v -run TestFmtRegex github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestBadDomainFilter github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestBadDomainAggregate github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestDNSTunnelFilter github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestDNSTunnelAggregate github.com/bearlyrunning/FindingTheNeedle/go/detection
go test -v -run TestBrowserSubProcDetection github.com/bearlyrunning/FindingTheNeedle/go/detection
```
