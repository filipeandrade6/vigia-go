#!/usr/bin/env bash

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

export PATH="$PATH:$(go env GOPATH)/bin"
protoc -I=$DIR/api/proto/v1 --go_out=. --go-grpc_out=. $DIR/api/proto/v1/*.proto