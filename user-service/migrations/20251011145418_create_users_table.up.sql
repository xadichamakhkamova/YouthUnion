CREATE TYPE gender AS ENUM ('MALE', 'FEMALE');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    identifier INTEGER UNIQUE NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    password_hash TEXT NOT NULL,
    faculty VARCHAR(100),
    course INTEGER,
    birth_date VARCHAR(20) NOT NULL,
    gender gender NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);
