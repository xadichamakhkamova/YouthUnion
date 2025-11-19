-- ENUM: sender type
CREATE TYPE sender_type AS ENUM (
    'UNKNOWN_SENDER',
    'ADMIN',
    'ORGANIZER'
);

-- ENUM: notification type
CREATE TYPE notification_type AS ENUM (
    'SYSTEM',
    'SCORE_UPDATE',
    'TEAM_INVITE',
    'EVENT_UPDATE',
    'GENERAL'
);

-- TABLE
CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    sender_id UUID NOT NULL,
    sender_type sender_type NOT NULL,

    user_id UUID,
    is_public BOOLEAN NOT NULL DEFAULT FALSE,

    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,

    type notification_type NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
