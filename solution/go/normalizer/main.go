package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	nlpb "github.com/bearlyrunning/FindingTheNeedle/solution/go/generated/normalizedlogpb"
)

const (
	dnsLog           = "../../data/dns/dns.log"
	dnsBinaryOutput  = "../../data/dns/dns_normalized.binpb"
	dnsJsonOutput    = "../../data/dns/dns_normalized.json"
	netLog           = "../../data/netflow/netflow.log"
	netBinaryOutput  = "../../data/netflow/netflow_normalized.binpb"
	netJsonOutput    = "../../data/netflow/netflow_normalized.json"
	execLog          = "../../data/execution/execution.log"
	execBinaryOutput = "../../data/execution/execution_normalized.binpb"
	execJsonOutput   = "../../data/execution/execution_normalized.json"
)

// Normalizer is an interface for log normalizers.
type Normalizer interface {
	getInput() string
	getBinaryOutput() string
	getJsonOutput() string
	normalize(string) *nlpb.NormalizedLog
}

// DNSNormalizer is the normalizer for DNS logs.
type DNSNormalizer struct {
	input        string
	binaryOutput string
	jsonOutput   string
}

// NetflowNormalizer is the normalizer for netflow logs.
type NetflowNormalizer struct {
	input        string
	binaryOutput string
	jsonOutput   string
}

// ExecutionNormalizer is the normalizer for execution logs.
type ExecutionNormalizer struct {
	input        string
	binaryOutput string
	jsonOutput   string
}

// Normalize reads from source log file and writes normalized logs to both wire and JSON formatted files.
func run(n Normalizer) error {
	// Set up reader.
	inputF, err := os.Open(n.getInput())
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err)
	}
	defer inputF.Close()
	s := bufio.NewScanner(inputF)

	// Set up binary / wire format output writer.
	binF, err := os.Create(n.getBinaryOutput())
	if err != nil {
		return fmt.Errorf("failed to create binary output file: %s", err)
	}
	defer binF.Close()
	binW := bufio.NewWriter(binF)

	// Set up JSON format output writer.
	jsonF, err := os.Create(n.getJsonOutput())
	if err != nil {
		return fmt.Errorf("failed to create json output file: %s", err)
	}
	defer jsonF.Close()
	jsonW := bufio.NewWriter(jsonF)

	for s.Scan() {
		// Normalizing line by line.
		msg := n.normalize(s.Text())

		// Encoding proto to wire format.
		en, err := proto.Marshal(msg)
		if err != nil {
			return fmt.Errorf("failed encoding log message %v: %v", msg, err)
		}

		// Write the size of the msg to the buffer.
		// This is necessary for streaming multiple messages.
		// See https://protobuf.dev/programming-guides/techniques/#streaming.
		size := make([]byte, 4)
		binary.LittleEndian.PutUint32(size, uint32(len(en)))
		binW.Write(size)

		// Write the msg itself to buffer.
		binW.Write(en)

		// Encoding proto to JSON format.
		j, err := protojson.Marshal(msg)
		if err != nil {
			return fmt.Errorf("failed converting log message to json format %v: %v", msg, err)
		}
		jsonW.Write(j)
		jsonW.WriteString("\n")
	}

	// Flush buffer.
	binW.Flush()
	jsonW.Flush()

	if err := s.Err(); err != nil {
		return fmt.Errorf("failed to scan file: %v", err)
	}
	return nil
}

func main() {
	normalizers := []Normalizer{
		&DNSNormalizer{dnsLog, dnsBinaryOutput, dnsJsonOutput},
		&NetflowNormalizer{netLog, netBinaryOutput, netJsonOutput},
		&ExecutionNormalizer{execLog, execBinaryOutput, execJsonOutput},
	}
	for _, n := range normalizers {
		if err := run(n); err != nil {
			log.Fatal(err)
		}

	}
}
