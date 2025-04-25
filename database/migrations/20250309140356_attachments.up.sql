CREATE TABLE attachments (
    id BIGSERIAL PRIMARY KEY,
    filename VARCHAR NOT NULL,
    filepath VARCHAR NOT NULL,
    filesize BIGINT NOT NULL,
    filetype VARCHAR NOT NULL,
    user_id BIGINT REFERENCES users(id) ON DELETE SET NULL,
    attachable_type VARCHAR NOT NULL,
    attachable_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_attachments_attachable ON attachments(attachable_type, attachable_id);
