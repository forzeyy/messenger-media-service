-- +migrate Up
CREATE TABLE IF NOT EXISTS metadata (
    id UUID PRIMARY KEY,
    owner_id TEXT,
    chat_id TEXT,
    type TEXT CHECK (type IN ('image', 'video', 'audio')),
    mime_type TEXT,
    size BIGINT,
    storage_path TEXT,
    created_at TIMESTAMPTZ
);