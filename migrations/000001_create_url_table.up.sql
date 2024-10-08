CREATE TABLE IF NOT EXISTS urls
(
    id           SERIAL PRIMARY KEY,
    short_url    VARCHAR(10) UNIQUE NOT NULL,
    original_url VARCHAR(255)       NOT NULL,
    created_at   TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);