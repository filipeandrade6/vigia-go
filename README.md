## VIGIA

#### ONDE PAREI

INSERT INTO usuarios (usuario_id, email, funcao, senha_hash) VALUES ('d03307d4-2b28-4c23-a004-3da25e5b8ce2', 'filipe.andrade@ssp.df.gov.br', '{ADMIN, MANAGER, USER}', '$2a$10$n531epIH68yygcV6sNNqluZtyPc3smWxbw1WoWDDhOIqUP1Py/GTq');

INSERT INTO cameras (camera_id, descricao, endereco_ip, porta, canal, usuario, senha, geolocalizacao) VALUES ('d03307d4-2b28-4c23-a004-3da25e5b8ce3', 'desc 1', '10.92.10.1', '1', '1', 'admin', 'admin', '-12.2332, -42.231');

Criando authenticação, ver link fixado de como add auth token no ctx da request do gRPC

ver como fica as dependencias/hierarquias, pois client depende do service e nao pode separar eles.

mover os service.go em internal/grpc/XXXX para internal/XXXX/grpc

mover os arquivos gerando para internal/api/pb/XXXXX

ativar metricas e demais antes de partir para os outros (processos/usuarios/servidores_de_gravacao/etc...)




quais erros devem ser devolvidos para o client...
entender o log




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
