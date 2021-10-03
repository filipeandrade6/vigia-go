CREATE TABLE cameras (
    camera_id UUID, -- TODO: POSTGRES tem DEFAULT gen_random_uuid()
    descricao TEXT NOT NULL,
    endereco_ip TEXT NOT NULL UNIQUE,
    porta INTEGER NOT NULL,
    canal INTEGER NOT NULL,
    usuario TEXT NOT NULL,
    senha TEXT NOT NULL,
    latitude TEXT NOT NULL,
    longitude TEXT NOT NULL,

    PRIMARY KEY (camera_id)
);