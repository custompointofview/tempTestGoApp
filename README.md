# TMP - App for Port Domain Service

## Requirements

To run the application you will need the following installed:

-   go@1.16.x
-   protobuf (for golang) @latest
-   docker@latest
-   docker-compose@latest

## Pre-execution & Code Changes

The app uses the Protobuf technology to generate relevant gRPC code.
If the _.proto_ files are changed, you should execute the `build/regenerate.sh` bash script.
This will regenerate all the needed _.go_ files for the application.

> Please do not modify the generated Go files.

> If you need a DB initialization, please add a ports.json file as in the above example. The database will be initialized at first run time.

## Architecture

There are 3 services:

-   `ClientAPI` - a service with a REST API dedicated to user interaction - port `:8000`
-   `PortDomainService` - a service with a gRPC interface that will interact with a Database to store diverse information resulted from user interaction - port `:8001`
-   `MongoDB` - a MongoDB Server dedicated to store the information - port `:27017`

## Docker & Docker Compose

The app uses `docker-compose` to create a working local environment.

### Build

    # run in terminal
    $ docker-compose build

> This will download & build the docker images needed for the application to function properly
