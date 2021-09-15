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

CREATE TABLE servidores_gravacao (
    servidor_gravacao_id UUID,
    endereco_ip INET NOT NULL,
    porta INTEGER NOT NULL,
    host TEXT NOT NULL,
    criado_em TIMESTAMP NOT NULL,
    editado_em TIMESTAMP,

    PRIMARY KEY (servidor_gravacao_id)
);

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

CREATE TABLE registros (
    registro_id UUID,
    placa TEXT NOT NULL,
    tipo_veiculo TEXT NOT NULL,
    cor_veiculo TEXT NOT NULL,
    marca_veiculo TEXT NOT NULL,
    armazenamento INT NOT NULL,
    processo_id UUID NOT NULL,
    criado_em TIMESTAMP NOT NULL,

    PRIMARY KEY (registro_id),
    FOREIGN KEY (processo_id) REFERENCES processos(processo_id) ON DELETE CASCADE
);

CREATE TABLE placas (
    placa_id UUID,
    placa TEXT NOT NULL,
    tipo_veiculo TEXT NOT NULL,
    cor_veiculo TEXT NOT NULL,
    marca_veiculo TEXT NOT NULL,
    criado_em TIMESTAMP NOT NULL,
    editado_em TIMESTAMP,

    PRIMARY KEY (placa_id)
);