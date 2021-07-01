#!/usr/bin/env bash

export PATH="$PATH:$(go env GOPATH)/bin"
protoc -I=./api/proto/v1 --go_out=. --go-grpc_out=. ./api/proto/v1/*.proto
#protoc -I=./internal/messages --go_out=. --go-grpc_out=. ./internal/messages/*.proto