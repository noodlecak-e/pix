CREATE TABLE pictures (
    id VARCHAR(36),
    event_id VARCHAR(36),
    user_id VARCHAR(36),
    image_base64 TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);