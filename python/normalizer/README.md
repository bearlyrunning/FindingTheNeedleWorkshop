# Log Normalization in Python

In this section, your task is, for each of the DNS, Netflow, and Execution normalizers:

1. Implement the `normalize(...)` function.
2. Complete the existing table-driven unit tests, and implement the remaining unimplemented tests.
3. Make sure your unit tests run as expected (see below for how to run these).

The locations you need to make code changes can be found by looking for the following comment string:

```
<TODO: Implement me!>
```

These changes need to take place in separate directories. We recommend completing this section in the following order:

1. `dns/dns.py`, with unit tests in `dns/test_dns.py`.
2. `netflow/netflow.py`, with unit tests in `netflow/test_netflow.py`.
3. `execution/execution.py`, with unit tests in `execution/test_execution.py`.

## Input and Output locations

The following table outlines where each normalizer ingests data from, and where its output is written to:

| **Normalizer** | **Code Location**      | **Input Data Location**            | **Output Data Location**                                                                                         |
|----------------|------------------------|------------------------------------|------------------------------------------------------------------------------------------------------------------|
| DNS            | `dns/dns.py`           | `../../data/dns/dns.log`           | Binary: `../../data/dns/dns_normalized.binpb` JSON: `../../data/dns/dns_normalized.json`                         |
| Execution      | `execution/execution.py` | `../../data/execution/execution.log` | Binary: `../../data/execution/execution_normalized.binpb` JSON: `../../data/execution/execution_normalized.json` |
| Netflow        | `netflow/netflow.py`     | `../../data/netflow/netflow.log`     | Binary: `../../data/netflow/netflow_normalized.binpb` JSON: `../../data/netflow/netflow_normalized.json`         |

## Running the normalizers

Run the following command in the local directory:

```
$ python3 main.py
```

## Running the unit tests

Run the following command in the local directory:

```
$ python3 -m unittest
```