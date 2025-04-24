CREATE TABLE templates
(
    id                   BIGSERIAL PRIMARY KEY,
    titulo               VARCHAR(255) NOT NULL,
    tabela           VARCHAR(255) NOT NULL,
    coluna_identificacao VARCHAR(255) NOT NULL,
    colunas              JSON         NOT NULL,
    created_at           TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP
);
