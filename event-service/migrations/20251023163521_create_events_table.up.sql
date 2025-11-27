CREATE TYPE event_type AS ENUM ('INDIVIDUAL', 'TEAM');

CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_type event_type NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    location VARCHAR(200),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, 
    max_participants INT NOT NULL,
    current_participants INT DEFAULT 0,
    min_team_size INT,
    max_team_size INT,
    status VARCHAR(20) DEFAULT 'UPCOMING' CHECK (status IN ('UPCOMING', 'ACTIVE', 'FINISHED', 'CANCELLED')),
    poster_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);