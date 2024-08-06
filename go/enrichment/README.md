# Enrichment in Golang

In this section, your task is, for each of the bad DNS resolution, DNS tunneling, and browser subprocess enrichment use cases:

1. Enrich your detections using our gRPC server, by implementing the `enrich()` function.
2. Parse the appropriate fields from the response, and add these to your `Signal` proto as appropriate.

The locations you need to make code changes can be found by looking for the following comment string:

```
<TODO: Implement me!>
```

These changes need to take place in separate files. We recommend completing this section in the following order (**note: the order is different to the previous sections!**):

1. `dns_tunnel.go`
2. `browser_sub_proc.go`
3. `bad_domain.go`

## Input and Output locations

The following table outlines where each enrichment module ingests data from, and where its output is written to:

| **Enrichment**                   | **Code Location**     | **Input Data Location**                   | **Output Data Location**                           |
|----------------------------------|-----------------------|-------------------------------------------|----------------------------------------------------|
| DNS tunnelling                   | `dns_tunnel.go`       | `../../data/signal/dns_tunnel.json`       | `../../data/signal/dns_tunnel.json_enriched`       |
| Suspicious browser child process | `browser_sub_proc.go` | `../../data/signal/browser_sub_proc.json` | `../../data/signal/browser_sub_proc.json_enriched` |
| Bad DNS domain resolution        | `bad_domain.go`       | `../../data/signal/bad_domain.json`       | `../../data/signal/bad_domain.json_enriched`       |

## Running your enrichment logic

To build and run the enrichment binary:

```
$ go build
$ ./enrichment --addr=findingtheneedle-wn5okteava-uc.a.run.app:443
```
