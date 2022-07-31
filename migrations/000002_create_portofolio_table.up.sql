-- Create tables
CREATE TABLE portofolio (
    id              BIGSERIAL       PRIMARY KEY,
    title           VARCHAR (255)   NOT NULL,
    descriptions    TEXT            NOT NULL,
    project_info    TEXT            NOT NULL,
    languages       VARCHAR (255)   NOT NULL,
    databases       VARCHAR (255)   NOT NULL,
    dates           VARCHAR (255)   NOT NULL,
    platform        VARCHAR (255)   NOT NULL,
    urls            VARCHAR (255)   NOT NULL,
    source_code     VARCHAR (255)   NOT NULL,
    created_at      TIMESTAMP       NOT NULL,
    updated_at      TIMESTAMP       NOT NULL
);