## VIGIA

COLOCAR DUAS PORTAS NO SERVIDOR DE GERENCIA
1 PARA INTERPROCESSOS COM GRAVACAO E OUTRO FRONTEND

#### ONDE PAREI

APRENDER BUFIO - SCANNER - IO.READ

NOVO BANCO COM HISTORICO DE NOTIFICAÇÕES

VER OS CASCADE

NÃO COMITAR O .ENV FILE PARA O VCS(GIT) AS CONFIG DEVE SER FORNECIDAS NA APLICACAO

HEALTHY CHECK - https://github.com/grpc/grpc/blob/master/doc/health-checking.md

NAO TEM GRPC PARA BROWSER https://github.com/grpc-ecosystem/grpc-gateway/blob/master/examples/internal/proto/examplepb/wrappers.proto

COMO O EVANS FAZ ITERATIVO - REPL

ADICIONAR AUTHORIZATION POR NIVEL DE SERVIÇO ASSIM COMO O TECH SCHOOL

COLOCAR require de na funcação do usuario ou criar if proibindo nulo	fmt.Println(fullMethodName)

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

### FEDORA

- Na instalação do protobuf - não instale com dnf install protoc - siga a respota... https://stackoverflow.com/questions/40025602/how-to-use-predifined-protobuf-type-i-e-google-protobuf-timestamp-proto-wit


NewUnit - cria um banco de dados de teste

NewIntegration - cria um db, alimenta ele  e constroi um autenticador (cria chave, cria um autenticador com essa chave)
Retorna um test { DB, LOG, AUTH, testing.T e função de teardown}

Token - gera um token autenticado para o usuario
store - usuarioStore, claims e token utilizando o test acima

como a verificação de auth fica na requisição da Store - não vou precisar testar o Authentication
