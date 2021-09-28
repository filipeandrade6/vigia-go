## VIGIA

#### ONDE PAREI

Add https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md ?

EU POSSO REGISTRAR OUTROS SERVICOS NO GRPCSERVER - isso permite separar os servicos em arquivos

VER O EVANS no TECH SCHOOL PARA CLIENT DE GRPC - acho que usa o grpc reflect...

CRIAR OS TESTE E FAZER TDD

#### TODO

- [ ] Colocar interface no querier
- [ ] Execução gRPC em contexto e Health Server https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869
- [ ] Verificar timezone na aplicação e quando abre o banco de dados

#### comandos *Makefile*:

- `make protobuf` gera os arquivos protobuf de acordo com os arquivos proto em /api/proto/v1
- `make run` executa as aplicações de gerencia e gravação
- `make test` executa os testes
- `make create-postgres` cria o container Docker de Postgres para desenvolvimento
- `make create-pgadmin` cria o container Docker de pgAdmin para desenvolvimento
- `make postgrse` reinicia o container Postgres
- `make pgadmin` reinicia o container pgAdmin

### boot

- docker-compose up informando qual o tipo de ambiente (dev, test, staging, prod)
- no docker-compose tem variaveis de ambiente que sobrescrevem o arquivo de configuracao
- carrega o arquivo de configuracao

### FEDORA

- Na instalação do protobuf - não instale com dnf install protoc - siga a respota... https://stackoverflow.com/questions/40025602/how-to-use-predifined-protobuf-type-i-e-google-protobuf-timestamp-proto-wit
