-- Migration script to create the RefreshToken table in PostgreSQL

CREATE TYPE refresh_token_status AS ENUM ('new', 'used');

CREATE TABLE IF NOT EXISTS refresh_tokens (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    jti VARCHAR NOT NULL,
    parent INTEGER REFERENCES refresh_tokens(id) ON DELETE SET NULL,
    status refresh_token_status NOT NULL DEFAULT 'new'
);

-- Indexes
CREATE INDEX idx_refresh_tokens_jti ON refresh_tokens(jti);
