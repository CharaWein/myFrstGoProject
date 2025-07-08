CREATE TABLE profiles (
    profile_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    friends TEXT[] NOT NULL DEFAULT '{}',
    subscribes_count INTEGER NOT NULL DEFAULT 0,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- Index for faster lookups by user_id
CREATE INDEX idx_profiles_user_id ON profiles(user_id);