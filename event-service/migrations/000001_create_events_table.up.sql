CREATE TYPE event_type AS ENUM ('INDIVIDUAL', 'TEAM');

CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_type event_type NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    location VARCHAR(255),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    created_by UUID NOT NULL, -- User Service dagi foydalanuvchi ID
    max_participants INT DEFAULT 0, -- 0 = cheklanmagan
    current_participants INT DEFAULT 0, -- qatnashayotganlar soni (realtime update)
    status VARCHAR(20) DEFAULT 'ACTIVE', -- ACTIVE, CANCELLED, FINISHED
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);
