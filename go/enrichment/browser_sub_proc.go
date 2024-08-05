package main

import (
	"context"

	enpb "github.com/bearlyrunning/FindingTheNeedle/go/generated/enrichmentpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/go/generated/signalpb"
)

func (bspe *BrowserSubProcEnricher) getName() string {
	return bspe.name
}

func (bspe *BrowserSubProcEnricher) getSignals() []*spb.Signal {
	return bspe.signals
}

func (bspe *BrowserSubProcEnricher) setSignals(sigs []*spb.Signal) {
	bspe.signals = sigs
}

// Skipping unit test as this involves setting up mock which is out of scope for this workshop.
func (bspe *BrowserSubProcEnricher) enrich(ctx context.Context, c enpb.EnrichmentClient) {
	// <TODO: Implement me!>
	// Populate the `source_ip` field in the spb.BrowserSubProc message for all signals.
	// Hint: Check the RPC methods supported by enpb.EnrichmentClient in the generated package file.
}
