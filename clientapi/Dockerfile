FROM golang:1.16.3-alpine3.12
ARG ROOT=/workspace

RUN apk --no-cache add git alpine-sdk bash

COPY . ${ROOT}

# make sure we have all the modules needed
WORKDIR ${ROOT}
RUN go mod download

# build service
WORKDIR ${ROOT}
RUN go build -v ./cmd/server
