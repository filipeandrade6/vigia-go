CREATE TABLE servidores_gravacao (
    servidor_gravacao_id UUID,
    endereco_ip INET NOT NULL,
    porta INTEGER NOT NULL,
    host TEXT NOT NULL,
    criado_em TIMESTAMP NOT NULL,
    editado_em TIMESTAMP,

    PRIMARY KEY (servidor_gravacao_id)
);