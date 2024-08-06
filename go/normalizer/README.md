# Log Normalization in Golang

In this section, your task is, for each of the DNS, Netflow, and Execution normalizers:

1. Implement the `normalize(...)` function.
2. Complete the existing table-driven unit tests, and implement the remaining unimplemented tests.
3. Make sure your unit tests run as expected (see below for how to run these).

The locations you need to make code changes can be found by looking for the following comment string:

```
<TODO: Implement me!>
```

These changes need to take place in separate files. We recommend completing this section in the following order:

1. `dns.go`, with unit tests in `dns_test.go`.
2. `netflow.go`, with unit tests in `netflow_test.go`.
3. `execution.go`, with unit tests in `execution_test.go`.

## Input and Output locations

The following table outlines where each normalizer ingests data from, and where its output is written to:

| **Normalizer** | **Code Location**      | **Input Data Location**            | **Output Data Location**                                                                                         |
|----------------|------------------------|------------------------------------|------------------------------------------------------------------------------------------------------------------|
| DNS            | `dns.go`           | `../../data/dns/dns.log`           | Binary: `../../data/dns/dns_normalized.binpb` JSON: `../../data/dns/dns_normalized.json`                         |
| Execution      | `execution.go` | `../../data/execution/execution.log` | Binary: `../../data/execution/execution_normalized.binpb` JSON: `../../data/execution/execution_normalized.json` |
| Netflow        | `netflow.go`     | `../../data/netflow/netflow.log`     | Binary: `../../data/netflow/netflow_normalized.binpb` JSON: `../../data/netflow/netflow_normalized.json`         |

## Running the normalizers

To build and run the normalizer binary:

```
$ go build
$ ./normalizer
```

## Running the unit tests

To run the unit tests:

```
go test -v -run TestDNSNormalizer github.com/bearlyrunning/FindingTheNeedle/go/normalizer
go test -v -run TestNetflowNormalizer github.com/bearlyrunning/FindingTheNeedle/go/normalizer
go test -v -run TestExecutionNormalizer github.com/bearlyrunning/FindingTheNeedle/go/normalizer
go test -v -run TestSplitWithEscape github.com/bearlyrunning/FindingTheNeedle/go/normalizer
```
