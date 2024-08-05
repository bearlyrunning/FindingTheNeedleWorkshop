package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	enpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/enrichmentpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
)

const (
	badDomainRuleName      = "bad_domain"
	dnsTunnelRuleName      = "dns_tunnel"
	browserSubProcRuleName = "browser_sub_proc"
	signalPath             = "../../data/signal/%s.json"
)

var (
	addr = flag.String("addr", "localhost:8080", "enrichment service address")
)

type Enricher interface {
	getName() string
	getSignals() []*spb.Signal
	setSignals(sigs []*spb.Signal)
	enrich(context.Context, enpb.EnrichmentClient)
}

type BadDomainEnricher struct {
	name    string
	signals []*spb.Signal
}

type DNSTunnelEnricher struct {
	name    string
	signals []*spb.Signal
}

type BrowserSubProcEnricher struct {
	name    string
	signals []*spb.Signal
}

func load(e Enricher) error {
	// Open signal file and set up scanner.
	path := fmt.Sprintf(signalPath, e.getName())
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %s", path, err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	var sigs []*spb.Signal
	for s.Scan() {
		sig := &spb.Signal{}
		if err := protojson.Unmarshal(s.Bytes(), sig); err != nil {
			return fmt.Errorf("failed converting json signal to proto %s: %v", s.Text(), err)
		}
		sigs = append(sigs, sig)
	}
	if err := s.Err(); err != nil {
		return fmt.Errorf("failed to scan file: %v", err)
	}
	e.setSignals(sigs)
	return nil
}

func output(e Enricher) error {
	// Set up output file for enriched signals.
	path := fmt.Sprintf(signalPath, e.getName()+"_enriched")
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %s", path, err)
	}
	defer out.Close()
	w := bufio.NewWriter(out)

	// Convert proto messages to Json format.
	for _, s := range e.getSignals() {
		j, err := protojson.Marshal(s)
		if err != nil {
			return fmt.Errorf("failed converting enriched signal message to json format %v: %v", s, err)
		}
		w.Write(j)
		w.WriteString("\n")
	}

	// Flush buffer.
	w.Flush()

	return nil
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	cred := insecure.NewCredentials()
	if *addr != "localhost:8080" {
		cred = credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	}
	opt := grpc.WithTransportCredentials(cred)
	conn, err := grpc.NewClient(*addr, opt)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Initiate a client.
	c := enpb.NewEnrichmentClient(conn)
	ctx := context.Background()

	// Set up and run enrichment.
	enrichers := []Enricher{
		&BadDomainEnricher{name: badDomainRuleName},
		&DNSTunnelEnricher{name: dnsTunnelRuleName},
		&BrowserSubProcEnricher{name: browserSubProcRuleName},
	}
	for _, e := range enrichers {
		if err := load(e); err != nil {
			log.Fatal(err)
		}
		e.enrich(ctx, c)
		if err := output(e); err != nil {
			log.Fatal(err)
		}
	}
}
