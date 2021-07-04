protobuf:
	bash scripts/protobuf-gen.sh

run:
	go run ./cmd/gerencia/main.go
	go run ./cmd/gravacao/main.go

test:
	# go clean --cache
	go test -cover ./...