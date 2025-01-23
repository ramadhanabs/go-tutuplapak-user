CREATE TABLE users (
    id SERIAL PRIMARY KEY,                     -- Auto-incrementing unique identifier
    email VARCHAR(255) UNIQUE DEFAULT '',      -- Email address (unique, optional)
    phone VARCHAR(20) UNIQUE DEFAULT '',       -- Phone number (unique, optional)
    password VARCHAR(255) NOT NULL,            -- Hashed password
    file_id VARCHAR(255) DEFAULT '',           -- Optional file ID
    file_uri TEXT DEFAULT '',                  -- Optional file URI
    file_thumbnail_uri TEXT DEFAULT '',        -- Optional thumbnail URI
    bank_account_name VARCHAR(255) DEFAULT '', -- Bank account name
    bank_account_holder VARCHAR(255) DEFAULT '', -- Bank account holder
    bank_account_number VARCHAR(50) DEFAULT '', -- Bank account number
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp of user creation
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp of last update
);

-- Add indexes for performance optimization
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_phone ON users (phone);
CREATE INDEX idx_users_file_id ON users (file_id);
CREATE INDEX idx_users_bank_account_number ON users (bank_account_number);