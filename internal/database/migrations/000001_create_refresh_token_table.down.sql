-- Migration script to drop the RefreshToken table and the enum type in PostgreSQL

-- Drop the index first
DROP INDEX IF EXISTS idx_refresh_tokens_jti;

-- Drop the table
DROP TABLE IF EXISTS refresh_tokens;

-- Drop the enum type
DROP TYPE IF EXISTS refresh_token_status;
