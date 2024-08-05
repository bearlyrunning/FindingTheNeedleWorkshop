package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	enpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/enrichmentpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
)

const (
	indicatorPath = "../../data/indicators/bad_domain.csv"
	columnNum     = 2
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

func loadIndicators() (map[string]enpb.Host_Platform, error) {
	hostToOS := make(map[string]enpb.Host_Platform)
	f, err := os.Open(indicatorPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %v", indicatorPath, err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		strs := strings.Split(s.Text(), ",")
		if len(strs) == columnNum {
			ioc, p := strs[0], strs[1]
			hostToOS[ioc] = enpb.Host_Platform(enpb.Host_Platform_value[p])
		}
	}
	if err = s.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %v", err)
	}
	return hostToOS, nil
}

// Skipping unit test as this involves setting up mock which is out of scope for this workshop.
func (bde *BadDomainEnricher) enrich(ctx context.Context, c enpb.EnrichmentClient) {
	var enrichedSigs []*spb.Signal
	hostToOS, err := loadIndicators()
	if err != nil {
		log.Print(err)
	}
	for _, sig := range bde.signals {
		resp, err := c.IPToHost(ctx, &enpb.IP{Ip: sig.GetBadDomain().GetSourceIp()})
		if err != nil {
			log.Printf("failed IPToHost conversion, output signals for manual investigation: %v", err)
			enrichedSigs = append(enrichedSigs, sig)
		}
		if resp.GetPlatform() != hostToOS[sig.GetBadDomain().GetBadDomain()] {
			log.Printf("domain indicator %s observed from %s - an operating system that isn't impacted", sig.GetBadDomain().GetBadDomain(), resp.GetPlatform())
		} else {
			sig.GetBadDomain().Hostname = resp.GetName()
			enrichedSigs = append(enrichedSigs, sig)
		}
		continue
	}
	bde.signals = enrichedSigs
}
