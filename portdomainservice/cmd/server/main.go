package main

import (
	"context"
	"log"
	"net"
	dbh "portdomainservice/dbhandlers"
	pds "portdomainservice/service"

	"google.golang.org/grpc"
)

const (
	MONGODB_ENDPOINT   = "mongodb://mongodbservice:27017"
	DB_DATABASE_NAME   = "portdomain"
	DB_COLLECTION_NAME = "ports"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("PORT DOMAIN SERVICE API v0.1 - Port Mapper")

	// create listener
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("Could not open listener on port 8001")
	}
	// crate db handler
	dbConfig := dbh.NewDBConfig(DB_DATABASE_NAME, DB_COLLECTION_NAME)
	dbHandler := dbh.NewDBHandler(dbh.DBHandler_MONGO, dbConfig)

	// create servers
	pdsServ := pds.NewServer(dbHandler)
	grpcServer := grpc.NewServer()

	// establish db connection
	if err := pdsServ.EstablishDBConnection(ctx, MONGODB_ENDPOINT); err != nil {
		log.Fatalf("DB Connection failed: %v", err)
	}

	// launch server
	log.Printf("Launching gRPC server on 8001...")
	pds.RegisterPortDomainServiceServer(grpcServer, pdsServ)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to launch server: %s", err)
	}
}
