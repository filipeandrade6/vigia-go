## VIGIA


#### ONDE PAREI

APLICAR WRAPPER E UNIQUE CONSTRAIN IGUAL USUARIO PARA OS OUTROS SERVICOS

HEALTHY CHECK - https://github.com/grpc/grpc/blob/master/doc/health-checking.md

NAO TEM GRPC PARA BROWSER https://github.com/grpc-ecosystem/grpc-gateway/blob/master/examples/internal/proto/examplepb/wrappers.proto

VERIFICAR OS ERRORS NO SERVICE

COMO O EVANS FAZ ITERATIVO - REPL

ADICIONAR AUTHORIZATION POR NIVEL DE SERVIÇO ASSIM COMO O TECH SCHOOL

ARRUMAR O SEED DOWN ESTA SEM OS TESTS PARA OS VALIDATORS

ADICIONAR TESTES PARA OS VALIDADORES E UNIQUE DO PG

Add https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md ?

EU POSSO REGISTRAR OUTROS SERVICOS NO GRPCSERVER - isso permite separar os servicos em arquivos

CRIAR OS TESTE E FAZER TDD

COLOCAR require de na funcação do usuario ou criar if proibindo nulo	fmt.Println(fullMethodName)

SEPARAR OS SERVIÇOS GRPC EM TIPO - CLIENT FRONT-END / GERENCIA INTER / GRAVACAO

ELE SEPAROU
 - ERROS NO DB DE database.ErrDBNotFound
 - ERROS de aplicação validate.ErrNotFound
 - ERROS viraram erro de validate - fora do doman de db

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
