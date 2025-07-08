CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    messages TEXT[] NOT NULL DEFAULT '{}'
);

-- Index for faster lookups by email
CREATE INDEX idx_users_email ON users(email);

-- Index for faster lookups by username
CREATE INDEX idx_users_username ON users(username);