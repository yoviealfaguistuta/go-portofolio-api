-- Create tables
CREATE TABLE about (
    id          BIGSERIAL       PRIMARY KEY,
    messages    TEXT            NOT NULL,
    created_at  TIMESTAMP       NOT NULL,
    updated_at  TIMESTAMP       NOT NULL
);

INSERT INTO about (id, messages, created_at, updated_at) VALUES 
(1, 'I am 23 years old and I am a web developer graduated in 2021. Coding on React JS, Laravel and Go Language with 2+ years experience. Serving 30+ users to build complex web application. My greatest strength is problem solving and logical thinking. I am very interested to use or learn the latest technology.', '2021-07-26 21:39:51', '2021-07-26 21:39:51');
