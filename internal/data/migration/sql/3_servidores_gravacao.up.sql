CREATE TABLE servidores_gravacao (
    servidor_gravacao_id UUID,
    endereco_ip TEXT NOT NULL,
    porta INTEGER NOT NULL,

    UNIQUE (endereco_ip, porta),
    PRIMARY KEY (servidor_gravacao_id)
);