CREATE TABLE IF NOT EXISTS team_invitations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    team_id UUID NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    inviter_id UUID NOT NULL,           -- kim taklif qildi (odatda leader)
    invited_user_id UUID NOT NULL,      -- kimga taklif yuborildi
    status VARCHAR(10) NOT NULL DEFAULT 'PENDING', -- PENDING / ACCEPTED / REJECTED
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    responded_at TIMESTAMP NULL,
    
    -- Unique constraint: bir userga bir team dan 1 ta invitation bo‘lishi
    UNIQUE(team_id, invited_user_id)
);

-- Indexlar query tezligi uchun
CREATE INDEX IF NOT EXISTS idx_team_invitations_team_id ON team_invitations(team_id);
CREATE INDEX IF NOT EXISTS idx_team_invitations_invited_user_id ON team_invitations(invited_user_id);
CREATE INDEX IF NOT EXISTS idx_team_invitations_status ON team_invitations(status);
