-- Create tables
CREATE TABLE about (
    id          BIGSERIAL       PRIMARY KEY,
    messages    TEXT            NOT NULL,
    created_at  TIMESTAMP       NOT NULL,
    updated_at  TIMESTAMP       NOT NULL
);

INSERT INTO about (id, messages, created_at, updated_at) VALUES 
(1, 'I am 23 years old and I am a software engineer with 1 years of experience. graduated in 2021 with a specialization in websites, mobile-based software, and video games. I prefer to spend time learning about new technologies related to programming.', '2021-07-26 21:39:51', '2021-07-26 21:39:51');
