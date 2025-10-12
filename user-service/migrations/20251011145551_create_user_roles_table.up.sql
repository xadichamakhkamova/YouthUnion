CREATE TABLE user_roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    role_id UUID NOT NULL REFERENCES roles_type(id),
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);
