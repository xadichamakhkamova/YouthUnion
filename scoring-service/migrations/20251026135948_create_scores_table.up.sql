CREATE TYPE scored_by_type AS ENUM ('ORGANIZER', 'ADMIN');

CREATE TABLE scores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    score_type VARCHAR(20) NOT NULL CHECK (score_type IN ('TEAM', 'INDIVIDUAL')),
    team_id UUID NULL REFERENCES teams(id) ON DELETE CASCADE,                 -- team konteksti
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,                 -- kimga baho berildi
    points INT NOT NULL CHECK (points >= 0),
    comment TEXT,
    scored_by_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,            -- kim baho berdi
    scored_by_type scored_by_type NOT NULL DEFAULT 'ORGANIZER',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);
