package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
	spb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/signalpb"
)

const (
	badDomainRuleName      = "bad_domain"
	dnsTunnelRuleName      = "dns_tunnel"
	browserSubProcRuleName = "browser_sub_proc"
	signalPath             = "../../data/signal/%s.json"
)

var (
	logPaths = []string{
		"../../data/dns/dns_normalized.binpb",
		"../../data/netflow/netflow_normalized.binpb",
		"../../data/execution/execution_normalized.binpb",
	}
)

type Detection interface {
	ruleName() string
	run() ([]*spb.Signal, error)
}

type NormalizedLog struct {
	dns       []*nlpb.DNS
	netflow   []*nlpb.Netflow
	execution []*nlpb.Execution
}

type BadDomainDetection struct {
	name string
	logs *NormalizedLog
	rr   *regexp.Regexp
}

type DNSTunnelDetection struct {
	name string
	logs *NormalizedLog
}

type BrowserSubProcDetection struct {
	name string
	logs *NormalizedLog
}

func (nl *NormalizedLog) load() error {
	for _, path := range logPaths {
		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %s", path, err)
		}
		defer f.Close()
		s := bufio.NewReader(f)

		for {
			sizeBuf := make([]byte, 4)
			_, err := io.ReadFull(s, sizeBuf)
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("failed fetching message size bytes: %v", err)
			}
			size := binary.LittleEndian.Uint32(sizeBuf)
			// log.Printf("msg size: %d\n", size)

			msgBuf := make([]byte, size)
			if _, err := io.ReadFull(s, msgBuf); err != nil {
				return fmt.Errorf("failed fetching message bytes: %v", err)
			}

			msg := &nlpb.NormalizedLog{}
			if err := proto.Unmarshal(msgBuf, msg); err != nil {
				return fmt.Errorf("failed decoding log message %v: %v", msg, err)
			}
			// log.Printf("msg retrieved: %v", msg)
			switch {
			case msg.GetDnsLog() != nil:
				nl.dns = append(nl.dns, msg.GetDnsLog())
			case msg.GetNetflowLog() != nil:
				nl.netflow = append(nl.netflow, msg.GetNetflowLog())
			case msg.GetExecutionLog() != nil:
				nl.execution = append(nl.execution, msg.GetExecutionLog())
			}
		}
	}
	// log.Printf("number of lines from dns %d netflow %d execution %d\n", len(nl.dns), len(nl.netflow), len(nl.execution))
	return nil
}

func output(name string, sigs []*spb.Signal) error {
	// Set up output file for signals.
	path := fmt.Sprintf(signalPath, name)
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %s", path, err)
	}
	defer out.Close()
	w := bufio.NewWriter(out)

	// Convert proto messages to Json format.
	for _, s := range sigs {
		j, err := protojson.Marshal(s)
		if err != nil {
			return fmt.Errorf("failed converting log message to json format %v: %v", s, err)
		}
		w.Write(j)
		w.WriteString("\n")
	}

	// Flush buffer.
	w.Flush()

	return nil
}

func main() {
	// Load all normalized logs in memory.
	nl := &NormalizedLog{}
	if err := nl.load(); err != nil {
		log.Fatalf("failed loading normalized logs: %v", err)
	}

	// Run detection rules.
	detections := []Detection{
		&BadDomainDetection{
			name: badDomainRuleName,
			logs: nl,
		},
		&DNSTunnelDetection{
			name: dnsTunnelRuleName,
			logs: nl,
		},
		&BrowserSubProcDetection{
			name: browserSubProcRuleName,
			logs: nl,
		},
	}
	for _, d := range detections {
		sigs, err := d.run()
		if err != nil {
			log.Fatal(err)
		}
		// Write any returned signals to disk.
		output(d.ruleName(), sigs)
	}
}
