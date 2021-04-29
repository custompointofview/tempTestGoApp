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

    {
        "AEAJM": {
            "name": "Ajman",
            "city": "Ajman",
            "country": "United Arab Emirates",
            "alias": [],
            "regions": [],
            "coordinates": [
            55.5136433,
            25.4052165
            ],
            "province": "Ajman",
            "timezone": "Asia/Dubai",
            "unlocs": [
            "AEAJM"
            ],
            "code": "52000"
        },
        "AEAUH": {
            "name": "Abu Dhabi",
            "coordinates": [
            54.37,
            24.47
            ],
            "city": "Abu Dhabi",
            "province": "Abu Z¸aby [Abu Dhabi]",
            "country": "United Arab Emirates",
            "alias": [],
            "regions": [],
            "timezone": "Asia/Dubai",
            "unlocs": [
            "AEAUH"
            ],
            "code": "52001"
        }
    }

> If you need a DB initialization, please add a ports.json file as in the above example. The database will be initialized at first run time.

Code & Test changes will be checked @ build time.
If the code is not formatted properly or the unit tests are failing the build will fail as well.

## Architecture

There are 3 services:

-   `ClientAPI` - a service with a REST API dedicated to user interaction - port `:8000`
-   `PortDomainService` - a service with a gRPC interface that will interact with a Database to store diverse information resulted from user interaction - port `:8001`
-   `MongoDB` - a MongoDB Server dedicated to store the information - port `:27017`

## Docker & Docker Compose

The app uses `docker-compose` to create a working local environment.

### Build

To build the application:

    # run in terminal
    $ docker-compose build

> This will download & build the docker images needed for the application to function properly

### Run

After building the docker images you are ready to execute the application.

    # run in terminal
    $ docker-compose up

### Interaction

After the services are up & running (it might take some time as the DB is initializing) you can interact with the application by using the ClientAPI.
This is available @ `localhost:8000`

The following commands are available:

    router.HandleFunc("/", handler.Homepage)
        -> will return a welcome message
    router.HandleFunc("/ports", handler.GetAllPorts).Methods("GET")
        -> not yet implemented
    router.HandleFunc("/port/{id}", handler.GetPort).Methods("GET")
        -> will return port with id
    router.HandleFunc("/port/{id}", handler.CreateOrUpdatePort).Methods("POST")
        -> will create or update port with id
        -> Content-Type - application/json
        -> same JSON as previously mentioned
    router.HandleFunc("/port/{id}", handler.DeletePort).Methods("DELETE")
        -> partially implemented

## Further notes

-   ClientAPI should open a stream to PortDomainService in order to send create & update commands
-   Everything should be thread safe (partially done right now)
-   Unit Testing should continue to a certain degree as the interaction between services can be mocked to have a proper look at local functionality
-   Functional testing can be employed - ex. testing values inserted by user with the values in DB
-   Performance testing has a role as well. This is not only related to the amount of data to initialize the DB with but also to the services interactions
