-- Create tables
CREATE TABLE certificates (
    id              BIGSERIAL       PRIMARY KEY,
    images          TEXT            NOT NULL,
    title           VARCHAR (255)   NOT NULL,
    publish         VARCHAR (100)   NOT NULL,
    credentials     VARCHAR (250)   NOT NULL DEFAULT '-',
    urls            VARCHAR (250)   NOT NULL DEFAULT '-',
    created_at      TIMESTAMP       NOT NULL,
    updated_at      TIMESTAMP       NOT NULL
);