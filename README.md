## VIGIA

SE DER ERRO DE DISCONNECT?

COLOCAR DUAS PORTAS NO SERVIDOR DE GERENCIA - 1 PARA INTERPROCESSOS COM GRAVACAO E OUTRO FRONTEND

TROCAR NOME HOUSEKEEPER POR HORAS

COLOCAR INTERFACE NAS COISAS

IMPLEMENTAR TESTES PARA GRAVACAO > PROCESSADOR

ver duplicata de armazenamento em registro, processo, servidor, etc... INDEX?

COLOCAR VERIFICADOR EXTERNO DE LOGIN E DISPONIBILIDADE...

#### ONDE PAREI

NOVO BANCO COM HISTORICO DE NOTIFICAÇÕES

VER OS CASCADE

CAMPOS AGREGADOS

NAO UTILIZAR TCP PARA SERVIDOR E GRAVACAO NO MESMO COMPUTADOR, USAR UNIX PIPE OU ALGO ASSIM TCP É BOM PARA REDES

NÃO COMITAR O .ENV FILE PARA O VCS(GIT) AS CONFIG DEVE SER FORNECIDAS NA APLICACAO

#### COR DO CARRO

EM_CAR_COLOR_WHITE        EM_CAR_COLOR_TYPE = iota // 白色
EM_CAR_COLOR_BLACK                                 // 黑色
EM_CAR_COLOR_RED                                   // 红色
EM_CAR_COLOR_YELLOW                                // 黄色
EM_CAR_COLOR_GRAY                                  // 灰色
EM_CAR_COLOR_BLUE                                  // 蓝色
EM_CAR_COLOR_GREEN                                 // 绿色
EM_CAR_COLOR_PINK                                  // 粉色
EM_CAR_COLOR_PURPLE                                // 紫色
EM_CAR_COLOR_DARK_PURPLE                           // 暗紫色
EM_CAR_COLOR_BROWN                                 // 棕色
EM_CAR_COLOR_MAROON                                // 粟色
EM_CAR_COLOR_SILVER_GRAY                           // 银灰色
EM_CAR_COLOR_DARK_GRAY                             // 暗灰色
EM_CAR_COLOR_WHITE_SMOKE                           // 白烟色
EM_CAR_COLOR_DEEP_ORANGE                           // 深橙色
EM_CAR_COLOR_LIGHT_ROSE                            // 浅玫瑰色
EM_CAR_COLOR_TOMATO_RED                            // 番茄红色
EM_CAR_COLOR_OLIVE                                 // 橄榄色
EM_CAR_COLOR_GOLDEN                                // 金色
EM_CAR_COLOR_DARK_OLIVE                            // 暗橄榄色
EM_CAR_COLOR_YELLOW_GREEN                          // 黄绿色
EM_CAR_COLOR_GREEN_YELLOW                          // 绿黄色
EM_CAR_COLOR_FOREST_GREEN                          // 森林绿
EM_CAR_COLOR_OCEAN_BLUE                            // 海洋绿
EM_CAR_COLOR_DEEP_SKYBLUE                          // 深天蓝
EM_CAR_COLOR_CYAN                                  // 青色
EM_CAR_COLOR_DEEP_BLUE                             // 深蓝色
EM_CAR_COLOR_DEEP_RED                              // 深红色
EM_CAR_COLOR_DEEP_GREEN                            // 深绿色
EM_CAR_COLOR_DEEP_YELLOW                           // 深黄色
EM_CAR_COLOR_DEEP_PINK                             // 深粉色
EM_CAR_COLOR_DEEP_PURPLE                           // 深紫色
EM_CAR_COLOR_DEEP_BROWN                            // 深棕色
EM_CAR_COLOR_DEEP_CYAN                             // 深青色
EM_CAR_COLOR_ORANGE                                // 橙色
EM_CAR_COLOR_DEEP_GOLDEN                           // 深金色
EM_CAR_COLOR_OTHER        = 255                    // 未识别、其他

#### TODO

- [ ] Colocar interface no querier
- [ ] Execução gRPC em contexto e Health Server https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869 - https://github.com/grpc/grpc/blob/master/doc/health-checking.md
- [ ] Verificar timezone na aplicação e quando abre o banco de dados
- [ ] Atualizar armazenamento mover imagens para novo local

#### comandos *Makefile*:

- `make protobuf` gera os arquivos protobuf de acordo com os arquivos proto em /api/proto/v1
- `make run` executa as aplicações de gerencia e gravação
- `make test` executa os testes
- `make create-postgres` cria o container Docker de Postgres para desenvolvimento
- `make create-pgadmin` cria o container Docker de pgAdmin para desenvolvimento
- `make postgres` reinicia o container Postgres
- `make pgadmin` reinicia o container pgAdmin

### FEDORA

- Na instalação do protobuf - não instale com dnf install protoc - siga a respota... https://stackoverflow.com/questions/40025602/how-to-use-predifined-protobuf-type-i-e-google-protobuf-timestamp-proto-wit


NewUnit - cria um banco de dados de teste

NewIntegration - cria um db, alimenta ele  e constroi um autenticador (cria chave, cria um autenticador com essa chave)
Retorna um test { DB, LOG, AUTH, testing.T e função de teardown}

Token - gera um token autenticado para o usuario
store - usuarioStore, claims e token utilizando o test acima

como a verificação de auth fica na requisição da Store - não vou precisar testar o Authentication
