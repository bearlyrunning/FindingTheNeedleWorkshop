// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/bearlyrunning/FindingTheNeedle/grpc/service"
)

var (
	ipToHost, hostToIP = make(map[string]string), make(map[string]string)
	hostToOS           = map[string]pb.Host_Platform{
		"hilversum": pb.Host_WINDOWS,
		"calydon":   pb.Host_MAC,
		"cyme":      pb.Host_WINDOWS,
	}
	hostnames = []string{
		"delphi",
		"bastion",
		"pydna",
		"pozzuoli",
		"foggia",
		"brest",
		"yokohama",
		"isfahan",
		"nijmegen",
		"haapajarvi",
		"cyme",
		"gaia",
		"demetrias",
		"pordenone",
		"hiroshima",
		"hermione",
		"iassus",
		"dallas",
		"crotone",
		"savona",
		"abydos",
		"ithaca",
		"modena",
		"stagirus",
		"geta",
		"syracuse",
		"braga",
		"genoa",
		"manfredonia",
		"cosenza",
		"apollonia",
		"gouda",
		"rhamnus",
		"lampsacus",
		"victoria",
		"hamedan",
		"civitavecchia",
		"helsinki",
		"ravenna",
		"aegina",
		"gallarate",
		"poway",
		"teramo",
		"pelkosenniemi",
		"reno",
		"montreal",
		"acharnae",
		"amsterdam",
		"messene",
		"padua",
		"suomussalmi",
		"piraeus",
		"viana",
		"meybod",
		"gorgan",
		"columbia",
		"pharae",
		"therma",
		"gortyn",
		"cagliari",
		"aprilia",
		"paphos",
		"tampa",
		"fasa",
		"thoricus",
		"masal",
		"leiden",
		"phocaea",
		"foligno",
		"rennes",
		"tegea",
		"osaka",
		"shanghai",
		"marseille",
		"calydon",
		"boston",
		"brindisi",
		"moncalieri",
		"ascoli",
		"ottawa",
		"sicyon",
		"troy",
		"marathon",
		"vittoria",
		"tours",
		"limassol",
		"oreus",
		"campobasso",
		"avellino",
		"rautavaara",
		"viareggio",
		"sanfrancisco",
		"heraclea",
		"olympia",
		"siouxcity",
		"barletta",
		"darab",
		"velletri",
		"andros",
		"lille",
		"boise",
		"adelaide",
		"taivassalo",
		"lasvegas",
		"milan",
		"naples",
		"pergamum",
		"magnesia",
		"sendai",
		"nicosia",
		"thebes",
		"bitonto",
		"valkeakoski",
		"washingtondc",
		"limoges",
		"parma",
		"fasham",
		"lisboa",
		"metz",
		"livorno",
		"tabriz",
		"minab",
		"perth",
		"amiens",
		"carrara",
		"salamis",
		"pteleum",
		"aenus",
		"santarosa",
		"tehran",
		"delft",
		"pandosia",
		"chioggia",
		"boulder",
		"faro",
		"apeldoorn",
		"rho",
		"thermum",
		"massa",
		"elis",
		"dubuque",
		"omaha",
		"potidaea",
		"alavieska",
		"hobart",
		"caltanissetta",
		"ialysus",
		"zacynthus",
		"sandiego",
		"nir",
		"haarlem",
		"acanthus",
		"cremona",
		"pisa",
		"hilversum",
		"scafati",
		"mytilene",
		"darwin",
		"angers",
		"trachis",
		"carpi",
		"sanantonio",
		"acireale",
		"utrecht",
		"bergamo",
		"santamonica",
		"naupactus",
		"brisbane",
		"acerra",
		"lecce",
		"catania",
		"halicarnassus",
		"oropus",
		"carystus",
		"nantes",
		"atlanta",
		"matera",
		"cnidus",
		"benevento",
		"baneh",
		"amphipolis",
		"thyreum",
		"venice",
		"nagasaki",
		"tokyo",
		"anzio",
		"lucca",
		"pylos",
		"nashville",
		"cythera",
		"eleusis",
		"trani",
		"catanzaro",
		"sansevero",
		"ahvaz",
		"modica",
		"nimes",
		"orimattila",
		"fiumicino",
		"losangeles",
		"miletus",
		"imola",
		"nice",
		"brescia",
		"arak",
		"honolulu",
		"methone",
		"aversa",
		"thespiae",
		"stlouis",
		"leeuwarden",
		"milwaukee",
		"sassari",
		"louisville",
		"bari",
		"larissa",
		"pagasae",
		"sapporo",
		"beijing",
		"caserta",
		"gela",
		"corinth",
		"phlius",
		"vicenza",
		"amlash",
		"udine",
		"rovigo",
		"kyoto",
		"pesaro",
		"athens",
		"ercolano",
		"kobe",
		"cleveland",
		"enontekio",
		"maastricht",
		"gortys",
		"megara",
		"alkmaar",
		"aegium",
		"collegno",
		"melbourne",
		"sparta",
		"laredo",
		"lepreum",
		"lemans",
		"pescara",
		"massilia",
		"pistoia",
		"saint",
		"denis",
		"bologna",
		"palermo",
		"indianapolis",
		"afragola",
		"placerville",
	}
)

type Server struct {
	pb.EnrichmentServer
}

func (s *Server) IPToHost(ctx context.Context, req *pb.IP) (*pb.Host, error) {
	if host, ok := ipToHost[req.GetIp()]; ok {
		h := &pb.Host{
			Name:     host,
			Platform: pb.Host_LINUX,
		}
		if p, ok := hostToOS[host]; ok {
			h.Platform = p
		}
		return h, nil
	}
	return nil, fmt.Errorf("no corresponding host found for IP address %s", req.GetIp())
}

func (s *Server) HostToIP(ctx context.Context, req *pb.Host) (*pb.IP, error) {
	if ip, ok := hostToIP[req.GetName()]; ok {
		return &pb.IP{Ip: ip}, nil
	}
	return nil, fmt.Errorf("no corresponding IP found for host %s", req.GetName())
}

func loadMaps() {
	i := 2
	for _, h := range hostnames {
		ipStr := fmt.Sprintf("10.20.30.%d", i)
		ipToHost[ipStr] = h
		hostToIP[h] = ipStr
		i++
	}
}

func main() {
	log.Printf("grpc-enrichment: starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	loadMaps()

	gs := grpc.NewServer()
	pb.RegisterEnrichmentServer(gs, &Server{})

	// Register reflection service on gRPC server.
	// See https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md.
	reflection.Register(gs)

	if err = gs.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
