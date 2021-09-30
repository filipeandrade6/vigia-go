INSERT INTO usuarios (usuario_id, email, funcao, senha_hash) VALUES
    ('d03307d4-2b28-4c23-a004-3da25e5b8ce2', 'filipe@andrade.com.br', '{ADMIN, MANAGER, USER}', '$2a$10$n531epIH68yygcV6sNNqluZtyPc3smWxbw1WoWDDhOIqUP1Py/GTq'),
    ('d03307d4-2b28-4c23-a004-3da25e5b8cf3', 'filipe@andrade2.com.br', '{ADMIN, MANAGER, USER}', '$2a$10$n531epIH68yygcV6sNNqluZtyPc3smWxbw1WoWDDhOIqUP1Py/GTq')
ON CONFLICT DO NOTHING;

INSERT INTO cameras (camera_id, descricao, endereco_ip, porta, canal, usuario, senha, geolocalizacao) VALUES
    ('d03307d4-2b28-4c23-a004-3da25e5b8ce3', 'desc 1', '10.20.30.40', '1', '1', 'admin', 'admin', '-12.2332, -42.231'),
    ('d03307d4-2b28-4c23-a004-3da25e5b8aa3', 'desc 2', '45.56.78.89', '1', '1', 'admin', 'admin', '-12.2332, -42.231')
ON CONFLICT DO NOTHING;

INSERT INTO servidores_gravacao (servidor_gravacao_id, endereco_ip, porta, host) VALUES
    ('d03307d4-2b28-4c23-a004-3da25e5b8bb1', '12.34.67.89', '6543', 'localhost'),
    ('d03307d4-2b28-4c23-a004-3da25e524bb1', '21.43.76.98', '3456', 'otherhost')
ON CONFLICT DO NOTHING;

INSERT INTO processos (processo_id, servidor_gravacao_id, camera_id, processador, adaptador, execucao) VALUES
    ('d03307d4-2b28-4c23-a004-3da32e5b8bb1', 'd03307d4-2b28-4c23-a004-3da25e5b8bb1', 'd03307d4-2b28-4c23-a004-3da25e5b8ce3', '1', '1', 'true'),
    ('d03307d4-2b28-4c23-a004-3da32e5b8a61', 'd03307d4-2b28-4c23-a004-3da25e5b8bb1', 'd03307d4-2b28-4c23-a004-3da25e5b8aa3', '1', '1', 'true')
ON CONFLICT DO NOTHING;