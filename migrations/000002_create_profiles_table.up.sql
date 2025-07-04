CREATE TABLE profiles (
    profile_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    friends TEXT[] DEFAULT '{}',
    subscribes_count INTEGER DEFAULT 0,
    user_id UUID UNIQUE NOT NULL REFERENCES users(user_id) ON DELETE CASCADE
);