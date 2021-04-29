package main

import (
	"log"
)

const (
	MONGODB_ENDPOINT   = "mongodb://mongodbservice:27017"
	DB_DATABASE_NAME   = "portdomain"
	DB_COLLECTION_NAME = "ports"
)

func main() {
	log.Println("PORT DOMAIN SERVICE API v0.1 - Port Mapper")
}
