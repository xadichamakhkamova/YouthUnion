CREATE TYPE score_target_type AS ENUM ('TEAM', 'INDIVIDUAL');
CREATE TYPE scored_by_type AS ENUM ('ORGANIZER', 'ADMIN');

CREATE TABLE scores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL,
    target_id UUID NOT NULL,                -- Team ID yoki User ID
    target_type score_target_type NOT NULL, -- TEAM yoki INDIVIDUAL
    points INT NOT NULL CHECK (points >= 0),
    comment TEXT,
    scored_by_id UUID NOT NULL,             -- kim qo‘ydi
    scored_by_type scored_by_type NOT NULL, -- ORGANIZER yoki ADMIN
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
