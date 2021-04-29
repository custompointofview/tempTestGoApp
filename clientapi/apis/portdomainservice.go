package apis

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	pds "clientapi/portdomainservice"

	"google.golang.org/grpc"
)

type PortDomainServAPI struct {
	ctx       context.Context
	endpoint  string
	conn      *grpc.ClientConn
	pdsClient pds.PortDomainServiceClient
}

func NewPortDomainServAPI(endpoint string) *PortDomainServAPI {
	return &PortDomainServAPI{
		endpoint: endpoint,
	}
}

func (p *PortDomainServAPI) Connect(ctx context.Context) error {
	p.ctx = ctx
	if p.endpoint == "" {
		log.Printf("Empty endpoint was provided. Skipping connection.")
		return nil
	}

	conn, err := GRPCConnect(ctx, p.endpoint)
	if err != nil {
		return err
	}
	p.conn = conn
	p.pdsClient = pds.NewPortDomainServiceClient(p.conn)
	return nil
}

func (p *PortDomainServAPI) Close() error {
	return p.conn.Close()
}

func (p *PortDomainServAPI) InitializeDatabase(jsonFilePath string) error {
	// parse file specified in path and send info to service
	// check if everything is OK with specified file
	f, err := os.Open(jsonFilePath)
	if err != nil {
		log.Printf("Error reading file=%v: %v", jsonFilePath, err.Error())
		return err
	}
	defer f.Close()

	if _, err := f.Stat(); err != nil {
		log.Printf("Could not obtain stat, handle error: %v", err.Error())
		return err
	}

	// create json decoder
	reader := bufio.NewReader(f)
	dec := json.NewDecoder(reader)

	i := 0
	// start reading file - begin with token
	if _, err := dec.Token(); err != nil {
		log.Printf("Error in reading initial token: %v", err)
		return err
	}
	// log.Printf("%T: %v\n", t, t)

	for dec.More() {
		// read portID - which should be unique
		portID, err := dec.Token()
		if err != nil {
			log.Printf("Error in reading token: %v", err)
			return err
		}
		// log.Printf("%T: %v\n", portID, portID)
		// unmarshal port details
		elm := &pds.Port{Id: fmt.Sprintf("%s", portID)}
		if err := dec.Decode(elm); err != nil {
			log.Printf("Error in decoding port: %v", err)
			return err
		}
		// log.Printf("%T: %v\n", elm, elm)

		// create or update port on the portdomainservice
		if err := p.CreateOrUpdatePort(elm); err != nil {
			return err
		}

		// sleep and count
		// time.Sleep(3 * time.Second)
		i++
	}
	if _, err = dec.Token(); err != nil {
		log.Printf("Error in reading final token: %v", err)
		return err
	}
	// log.Printf("%T: %v\n", t, t)

	log.Printf("Total of [%v] objects created.\n", i)
	return nil
}

func (p *PortDomainServAPI) GetPort(id string) (*pds.Port, error) {
	req := &pds.PortRequest{
		Port: &pds.Port{
			Id: id,
		},
	}
	response, err := p.pdsClient.GetPort(p.ctx, req)
	if err != nil {
		return nil, err
	}
	log.Printf("Got Port: %v", response.Port)
	return response.Port, nil
}
func (p *PortDomainServAPI) GetAllPorts() (map[string]*pds.Port, error) {
	// get all ports from service
	return nil, nil
}
func (p *PortDomainServAPI) CreateOrUpdatePort(port *pds.Port) error {
	// create/update port on service
	req := &pds.PortRequest{
		Port: port,
	}
	response, err := p.pdsClient.CreateOrUpdatePort(p.ctx, req)
	if err != nil {
		return err
	}
	log.Printf("Created or Updated Port: %v", response)
	return nil
}
func (p *PortDomainServAPI) DeletePort(id string) error {
	// delete port from service
	return nil
}
