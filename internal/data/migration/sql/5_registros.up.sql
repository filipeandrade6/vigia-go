CREATE TABLE registros (
    registro_id UUID,
    placa TEXT NOT NULL,
    tipo_veiculo TEXT NOT NULL,
    cor_veiculo TEXT NOT NULL,
    marca_veiculo TEXT NOT NULL,
    armazenamento INT NOT NULL,
    processo_id UUID NOT NULL,

    PRIMARY KEY (registro_id),
    FOREIGN KEY (processo_id) REFERENCES processos(processo_id) ON DELETE CASCADE
);