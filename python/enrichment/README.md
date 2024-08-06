# Enrichment in Python

In this section, your task is, for each of the bad DNS resolution, DNS tunneling, and browser subprocess enrichment use cases:

1. Enrich your detections using our gRPC server, by implementing the `enrich()` function.
2. Parse the appropriate fields from the response, and add these to your `Signal` proto as appropriate.

The locations you need to make code changes can be found by looking for the following comment string:

```
<TODO: Implement me!>
```

These changes need to take place in separate files. We recommend completing this section in the following order (**note: the order is different to the previous sections!**):

1. `dns_tunnel.py`
2. `browser_sub_proc.py`
3. `bad_domain.py`

## Input and Output locations

The following table outlines where each enrichment module ingests data from, and where its output is written to:

| **Enrichment**                   | **Code Location**     | **Input Data Location**                   | **Output Data Location**                           |
|----------------------------------|-----------------------|-------------------------------------------|----------------------------------------------------|
| DNS tunnelling                   | `dns_tunnel.py`       | `../../data/signal/dns_tunnel.json`       | `../../data/signal/dns_tunnel.json_enriched`       |
| Suspicious browser child process | `browser_sub_proc.py` | `../../data/signal/browser_sub_proc.json` | `../../data/signal/browser_sub_proc.json_enriched` |
| Bad DNS domain resolution        | `bad_domain.py`       | `../../data/signal/bad_domain.json`       | `../../data/signal/bad_domain.json_enriched`       |

## Running your enrichment logic

Run the following command in the local directory:

```
$ python3 main.py findingtheneedle-wn5okteava-uc.a.run.app:443
```
