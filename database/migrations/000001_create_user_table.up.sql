CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users(
    internal_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password text NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    public_id UUID NOT NULL DEFAULT gen_random_uuid(),
    CONSTRAINT user_public_id_unique UNIQUE (public_id) 
)