package main

import (
	"context"
	"log"

	enpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/enrichmentpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
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
	for _, sig := range bspe.signals {
		resp, err := c.HostToIP(ctx, &enpb.Host{Name: sig.GetBrowserSubProc().GetExecution().GetHostname()})
		if err != nil {
			log.Printf("failed HostToIP conversion: %v", err)
		} else {
			sig.GetBrowserSubProc().SourceIp = resp.GetIp()
		}
	}
}
