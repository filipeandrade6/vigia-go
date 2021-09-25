CREATE TABLE placas (
    placa_id UUID,
    placa TEXT NOT NULL,
    tipo_veiculo TEXT NOT NULL,
    cor_veiculo TEXT NOT NULL,
    marca_veiculo TEXT NOT NULL,

    PRIMARY KEY (placa_id)
);