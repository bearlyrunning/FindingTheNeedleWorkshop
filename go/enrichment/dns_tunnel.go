package main

import (
	"context"

	enpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/enrichmentpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/go/generated/signalpb"
)

func (dte *DNSTunnelEnricher) getName() string {
	return dte.name
}

func (dte *DNSTunnelEnricher) getSignals() []*spb.Signal {
	return dte.signals
}

func (dte *DNSTunnelEnricher) setSignals(sigs []*spb.Signal) {
	dte.signals = sigs
}

// Skipping unit test as this involves setting up mock which is out of scope for this workshop.
func (dte *DNSTunnelEnricher) enrich(ctx context.Context, c enpb.EnrichmentClient) {
	// <TODO: Implement me!>
	// Populate the `hostname` field in the spb.DNSTunnel message for all signals.
	// Hint: Check the RPC methods supported by enpb.EnrichmentClient in the generated package file.
}
