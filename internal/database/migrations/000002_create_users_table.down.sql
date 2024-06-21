-- Migration script to drop the User table in PostgreSQL

-- Drop the indexes first
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_email;

-- Drop the table
DROP TABLE IF EXISTS users;
