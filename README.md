## VIGIA

#### comandos *Makefile*:

- `make protobuf`gera os arquivos protobuf de acordo com os arquivos proto em /api/proto/v1
- `make run` executa as aplicações de gerencia e gravação
- `make test` executa os testes
- `make create-postgres` cria o container Docker de Postgres para desenvolvimento
- `make create-pgadmin`cria o container Docker de pgAdmin para desenvolvimento
- `make postgrse` reinicia o container Postgres
- `make pgadmin` reinicia o container pgAdmin

### boot

- docker-compose up informando qual o tipo de ambiente (dev, test, staging, prod)
- no docker-compose tem variaveis de ambiente que sobrescrevem o arquivo de configuracao
- carrega o arquivo de configuracao
