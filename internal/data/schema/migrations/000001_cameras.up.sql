CREATE TABLE IF NOT EXISTS cameras (
    id serial PRIMARY KEY,
    descricao VARCHAR(100) NOT NULL,
    ip VARCHAR(39) NOT NULL,
    porta SMALLINT NOT NULL,
    canal VARCHAR(30) NOT NULL,
    usuario_camera VARCHAR(30) NOT NULL,
    senha_camera VARCHAR(30) NOT NULL,
    cidade VARCHAR(30) NOT NULL,
    geolocalizacao VARCHAR(50) NOT NULL,
    marca VARCHAR(30) NOT NULL,
    modelo VARCHAR(30) NOT NULL,
    informacao VARCHAR(300) NOT NULL,
    editado_em TIMESTAMP NOT NULL,
);