#!/bin/bash

# generating code for portdomainservice
protoc --proto_path=../portdomainservice/protos/ --go_out=../portdomainservice/service --go_out=../clientapi/portdomainservice --go-grpc_out=../portdomainservice/service --go-grpc_out=../clientapi/portdomainservice --go_opt=paths=source_relative ../portdomainservice/protos/portdomainservice.proto

