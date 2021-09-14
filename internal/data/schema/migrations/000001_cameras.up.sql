CREATE TABLE cameras (
    camera_id UUID,
    descricao TEXT NOT NULL,
    endereco_ip TEXT NOT NULL,
    porta INTEGER NOT NULL,
    canal INTEGER NOT NULL,
    usuario TEXT NOT NULL,
    senha TEXT NOT NULL,
    geolocalizacao TEXT NOT NULL,
    criado_em TIMESTAMP NOT NULL,
    editado_em TIMESTAMP,

    PRIMARY KEY (camera_id)
);