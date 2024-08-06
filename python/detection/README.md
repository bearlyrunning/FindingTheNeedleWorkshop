# Detection Logic in Python

In this section, your task is, for each of the bad DNS resolution, DNS tunneling, and browser subprocess detections:

1. Design the detection logic for each detection rule, and implement the `run()` function.
2. Implement the unimplemented unit tests.
3. Make sure your unit tests run as expected.

The locations you need to make code changes can be found by looking for the following comment string:

```
<TODO: Implement me!>
```

These changes need to take place in separate files. We recommend completing this section in the following order:

1. `bad_domain.py`, with unit tests in `test_bad_domain.py`.
2. `dns_tunnel.py`, with unit tests in `test_dns_tunnel.py`.
3. `browser_sub_proc.py`, with unit tests in `test_browser_sub_proc.py`.

## Input and Output locations

The following table outlines where each detection rule ingests data from, and where its output is written to:

| **Rule**                         | **Code Location**     | **Input Data Location**                           | **Output Data Location**                  |
|----------------------------------|-----------------------|---------------------------------------------------|-------------------------------------------|
| Bad DNS domain resolution        | `bad_domain.py`       | `../../data/dns/dns_normalized.binpb"`            | `../../data/signal/bad_domain.json`       |
| DNS tunnelling                   | `dns_tunnel.py`       | `../../data/netflow/netflow_normalized.binpb`     | `../../data/signal/dns_tunnel.json`       |
| Suspicious browser child process | `browser_sub_proc.py` | `../../data/execution/execution_normalized.binpb` | `../../data/signal/browser_sub_proc.json` |

## Running your detection logic

Run the following command in the local directory:

```
$ python3 main.py
```

## Running the unit tests

Run the following command in the local directory:

```
$ python3 -m unittest
```