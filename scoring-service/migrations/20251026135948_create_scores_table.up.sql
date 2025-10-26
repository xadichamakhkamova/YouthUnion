CREATE TYPE scored_by_type AS ENUM ('ORGANIZER', 'ADMIN');

CREATE TABLE scores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL,
    team_id UUID NOT NULL,                 -- team konteksti
    user_id UUID NOT NULL,                 -- kimga baho berildi
    points INT NOT NULL CHECK (points >= 0),
    comment TEXT,
    scored_by_id UUID NOT NULL,            -- kim baho berdi
    scored_by_type scored_by_type NOT NULL DEFAULT 'ORGANIZER',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (event_id, team_id, user_id, scored_by_id)
);
