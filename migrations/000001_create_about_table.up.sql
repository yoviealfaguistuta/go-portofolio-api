-- Create tables
CREATE TABLE about (
    id          BIGSERIAL       PRIMARY KEY,
    messages    TEXT            NOT NULL,
    created_at  TIMESTAMP       NOT NULL,
    updated_at  TIMESTAMP       NOT NULL
);