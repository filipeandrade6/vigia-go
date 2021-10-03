CREATE TABLE veiculos (
    veiculo_id UUID,
    placa TEXT NOT NULL UNIQUE,
    tipo TEXT NOT NULL,
    cor TEXT NOT NULL,
    marca TEXT NOT NULL,
    info TEXT NOT NULL,

    UNIQUE (placa),
    PRIMARY KEY (veiculo_id)
);