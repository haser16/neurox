CREATE SCHEMA IF NOT EXISTS neurox;

CREATE TABLE neurox.users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE CHECK(char_length(username) BETWEEN 2 AND 100),
    email VARCHAR(100) NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    phone VARCHAR(15) CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    ),
    avatar VARCHAR(200),
    password CHAR(150) NOT NULL
);

CREATE TABLE neurox.requests(
    id SERIAL PRIMARY KEY,
    prompt VARCHAR(1000) NOT NULL,
    image VARCHAR(200),
    user_id INT NOT NULL REFERENCES pollify.users(id) ON DELETE CASCADE
);