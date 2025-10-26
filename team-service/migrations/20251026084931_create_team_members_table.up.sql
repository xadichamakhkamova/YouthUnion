CREATE TABLE team_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    team_id UUID NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    user_id UUID NOT NULL,                 -- UserService dagi foydalanuvchi ID
    role VARCHAR(20) NOT NULL DEFAULT 'MEMBER',  -- LEADER yoki MEMBER
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
