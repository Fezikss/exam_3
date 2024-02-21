CREATE TABLE books(
    id uuid NOT NULL PRIMARY KEY,
    name VARCHAR(30),
    author_name VARCHAR(30),
    page_number INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);