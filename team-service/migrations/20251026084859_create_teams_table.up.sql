CREATE TABLE teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    leader_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,          
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,           
    is_ready BOOLEAN DEFAULT FALSE,   -- jamoa to‘liq tayyor bo‘ldimi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
