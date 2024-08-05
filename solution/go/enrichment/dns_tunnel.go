package main

import (
	"context"
	"log"

	enpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/enrichmentpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
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
	for _, sig := range dte.signals {
		resp, err := c.IPToHost(ctx, &enpb.IP{Ip: sig.GetDnsTunnel().SourceIp})
		if err != nil {
			log.Printf("failed IPToHost conversion: %v", err)
		} else {
			sig.GetDnsTunnel().Hostname = resp.GetName()
		}
	}
}
