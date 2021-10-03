CREATE TABLE registros (
    registro_id UUID,
    processo_id UUID NOT NULL,
    placa TEXT NOT NULL,
    tipo_veiculo TEXT NOT NULL,
    cor_veiculo TEXT NOT NULL,
    marca_veiculo TEXT NOT NULL,
    armazenamento TEXT NOT NULL,
    confianca DECIMAL NOT NULL,
    criado_em TIMESTAMP WITH TIME ZONE NOT NULL,


    PRIMARY KEY (registro_id),
    FOREIGN KEY (processo_id) REFERENCES processos(processo_id) ON DELETE NO ACTION
);