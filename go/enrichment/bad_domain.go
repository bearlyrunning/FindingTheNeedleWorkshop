package main

import (
	"context"

	enpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/enrichmentpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/go/generated/signalpb"
)

const (
	indicatorPath = "../../data/indicators/bad_domain.csv"
)

func (bde *BadDomainEnricher) getName() string {
	return bde.name
}

func (bde *BadDomainEnricher) getSignals() []*spb.Signal {
	return bde.signals
}

func (bde *BadDomainEnricher) setSignals(sigs []*spb.Signal) {
	bde.signals = sigs
}

// Skipping unit test as this involves setting up mock which is out of scope for this workshop.
func (bde *BadDomainEnricher) enrich(ctx context.Context, c enpb.EnrichmentClient) {
	// <TODO: Implement me!>
	// NOTE: This enrichment is slightly more complex compared to `enrichment/browser_sub_proc.go`
	//       and `enrichment/dns_tunnel.go`. Please feel free to complete them in whatever order you prefer.

	// The bad_domain.csv file contains two fields:
	// - a domain indicator.
	// - the OS this indicator is applicable to.
	// Let's check if the source / client IP of the bad DNS traffic can be linked to a host with an applicable OS.
	// - If the traffic is observed from an irrelevant OS, drop the signal.
	// - If the host matches the relevant OS, populate the `hostname` field in the spb.BadDomain message for the signal.
	// And then return the enriched signals.

	// Hint #1: Check the RPC methods supported by enpb.EnrichmentClient in the generated package file.
	// Hint #2: What data structure can be used to allow easy lookup of the file content at `indicatorPath`?

	var enrichedSigs []*spb.Signal

	for _, sig := range bde.signals {
		// <TODO: Implement me!>
	}
	bde.signals = enrichedSigs
}
