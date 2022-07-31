-- Create tables
CREATE TABLE p_images (
    id              BIGSERIAL  PRIMARY KEY,
    id_portfolio    BIGINT     NOT NULL,
    images          TEXT       NOT NULL,
    orders          INT        NOT NULL,
    created_at      TIMESTAMP  NOT NULL,
    updated_at      TIMESTAMP  NOT NULL,

    CONSTRAINT fk_p_images_id_portfolio FOREIGN KEY("id_portfolio") REFERENCES "portofolio"("id") ON DELETE CASCADE
);