CREATE TABLE IF NOT EXISTS "user_credential"
(
    -- column
    id         VARCHAR PRIMARY KEY NOT NULL,
    username   VARCHAR UNIQUE      NOT NULL,
    email      VARCHAR UNIQUE      NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);