CREATE TABLE usuarios (
    usuario_id UUID,
    email TEXT UNIQUE,
    funcao TEXT[],
    senha_hash TEXT,

    PRIMARY KEY (usuario_id)
);