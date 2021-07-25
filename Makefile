protobuf:
	bash scripts/protobuf-gen.sh

run:
	go run ./cmd/gerencia/main.go
	go run ./cmd/gravacao/main.go

test:
	# go clean --cache
	go test -cover ./...

create-postgres:
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=secret -v ${PWD}/dev/postgresql/:/var/lib/postgresql/data --name dev-postgres postgres

create-pgadmin:
	docker run -d -p 80:80 -e 'PGADMIN_DEFAULT_EMAIL=user@domain.local' -e 'PGADMIN_DEFAULT_PASSWORD=secret' --name dev-pgadmin dpage/pgadmin4

postgres:
	docker restart dev-postgres

pgadmin:
	docker restart dev-pgadmin

docker-stop:
	docker stop dev-pgadmin dev-postgres