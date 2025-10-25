CREATE TABLE IF NOT EXISTS event_participants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    user_id UUID NOT NULL, -- User Service dagi foydalanuvchi ID
    role VARCHAR(20) DEFAULT 'ATTENDEE', -- ATTENDEE, ORGANIZER, SPEAKER
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);
