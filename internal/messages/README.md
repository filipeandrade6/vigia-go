export PATH="$PATH:$(go env GOPATH)/bin"

protoc -I=./internal/messages --go_out=. --go-grpc_out=. ./internal/messages/*.proto