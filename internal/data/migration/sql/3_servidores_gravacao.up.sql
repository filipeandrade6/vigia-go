CREATE TABLE servidores_gravacao (
    servidor_gravacao_id UUID,
    endereco_ip TEXT NOT NULL,
    porta INTEGER NOT NULL,
    host TEXT NOT NULL,

    PRIMARY KEY (servidor_gravacao_id)
);