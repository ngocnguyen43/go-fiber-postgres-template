-- Migration script to create the RefreshToken table in PostgreSQL

-- Create the ENUM type if it does not exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'refresh_token_status') THEN
        CREATE TYPE refresh_token_status AS ENUM ('new', 'used');
    END IF;
END $$;

-- Create the table if it does not exist
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    jti VARCHAR NOT NULL,
    parent INTEGER REFERENCES refresh_tokens(id) ON DELETE SET NULL,
    status refresh_token_status NOT NULL DEFAULT 'new'
);

-- Create the index if it does not exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_class WHERE relname = 'idx_refresh_tokens_jti') THEN
        CREATE INDEX idx_refresh_tokens_jti ON refresh_tokens(jti);
    END IF;
END $$;
