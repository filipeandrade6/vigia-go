CREATE TABLE processos (
    processo_id UUID,
    servidor_gravacao_id UUID,
    camera_id UUID,
    processador SMALLINT NOT NULL,
    adaptador SMALLINT NOT NULL,
    execucao BOOLEAN NOT NULL,
    criado_em TIMESTAMP NOT NULL,
    editado_em TIMESTAMP,

    PRIMARY KEY (processo_id),
    FOREIGN KEY (servidor_gravacao_id) REFERENCES servidores_gravacao(servidor_gravacao_id) ON DELETE CASCADE
);