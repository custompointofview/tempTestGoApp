package main

import (
	"flag"
	"log"
)

// DEFAULT PORTDOMAINSERVICE ENDPOINT
const PORTDOMAINSERVICE_ENDPOINT = "portdomainservice:8001"

func main() {
	log.Println("CLIENT REST API v0.1 - Port Router")
	portsFilePath := flag.String("ports", "", "Path to ports file")
	flag.Parse()

	log.Println("PortsFile:", portsFilePath)
}
