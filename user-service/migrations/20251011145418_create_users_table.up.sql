CREATE TYPE gender_enum AS ENUM ('MALE', 'FEMALE');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    identifier INT NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    password_hash TEXT NOT NULL,
    faculty VARCHAR(150),
    course SMALLINT,
    birth_date VARCHAR(20) NOT NULL,
    gender gender_enum NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);
