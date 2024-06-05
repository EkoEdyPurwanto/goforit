CREATE TABLE IF NOT EXISTS "user_credential"
(
    -- column
    id         VARCHAR PRIMARY KEY NOT NULL,
    username   VARCHAR(15) UNIQUE  NOT NULL,
    email      VARCHAR(50) UNIQUE  NOT NULL,
    password   VARCHAR(100)        NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);