## VIGIA

#### ONDE PAREI

refatorando

#### TODO

- [ ] Colocar interface no querier
- [ ] Aplicação executando em Docker
- [ ] Um único Dockerfile para ambos? (https://docs.docker.com/develop/develop-images/multistage-build/#use-an-external-image-as-a-stage)
- [ ] https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e
- [ ] Multi-stage imagem compartilhada entre Gerencia e Gravação
- [ ] Testes executando (iofs deprecated em embed utilizado go-migrations)
- [ ] Como fazer update de sistema
- [ ] DevOps
- [ ] Formatar arquivo de Dockerfile
- [ ] Label build
- [ ] Execução gRPC em contexto e Health Server https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869

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
