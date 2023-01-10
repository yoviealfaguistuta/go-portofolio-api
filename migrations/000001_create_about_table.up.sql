-- Create tables
CREATE TABLE about (
    id          BIGSERIAL       PRIMARY KEY,
    messages    TEXT            NOT NULL,
    created_at  TIMESTAMP       NOT NULL,
    updated_at  TIMESTAMP       NOT NULL
);

INSERT INTO about (id, messages, created_at, updated_at) VALUES 
(1, 'I am 23 years old and I am a web developer who graduated in 2021. Coding on React JS and Laravel with 2+ years of experience. I Already serve several clients to build web based applications. My greatest strength in technical terms is problem solving and logical thinking. I am also very enthusiastic and interested to use or learn new things.', '2021-07-26 21:39:51', '2021-07-26 21:39:51');
