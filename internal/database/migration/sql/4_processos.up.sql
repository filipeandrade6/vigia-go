CREATE TABLE processos (
    processo_id UUID,
    servidor_gravacao_id UUID,
    camera_id UUID,
    processador SMALLINT NOT NULL,
    adaptador SMALLINT NOT NULL,

    PRIMARY KEY (processo_id),
    FOREIGN KEY (servidor_gravacao_id) REFERENCES servidores_gravacao(servidor_gravacao_id) ON DELETE CASCADE
);