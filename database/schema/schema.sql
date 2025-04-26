CREATE TABLE users
(
    id       BIGSERIAL PRIMARY KEY,
    name     TEXT NOT NULL,
    email    TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE clients
(
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    email      TEXT NOT NULL,
    phone      TEXT,
    address    TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE projects
(
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT,
    client_id   BIGINT REFERENCES clients (id) ON DELETE CASCADE,
    status      TEXT NOT NULL DEFAULT 'pending',
    start_date  DATE,
    end_date    DATE,
    created_at  TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP     DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks
(
    id           BIGSERIAL PRIMARY KEY,
    title        TEXT   NOT NULL,
    description  TEXT,
    project_id   BIGINT REFERENCES projects (id) ON DELETE CASCADE,
    assigned_to  BIGINT REFERENCES users (id) ON DELETE SET NULL,
    status       TEXT   NOT NULL DEFAULT 'pending',
    priority     TEXT   NOT NULL DEFAULT 'medium',
    due_date     DATE,
    completed_at TIMESTAMP,
    created_at   TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE subtasks
(
    id           BIGSERIAL PRIMARY KEY,
    title        TEXT   NOT NULL,
    description  TEXT,
    task_id      BIGINT REFERENCES tasks (id) ON DELETE CASCADE,
    assigned_to  BIGINT REFERENCES users (id) ON DELETE SET NULL,
    status       TEXT   NOT NULL DEFAULT 'pending',
    due_date     DATE,
    completed_at TIMESTAMP,
    created_at   TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments
(
    id               BIGSERIAL PRIMARY KEY,
    content          TEXT   NOT NULL,
    user_id          BIGINT REFERENCES users (id) ON DELETE SET NULL,
    commentable_type TEXT   NOT NULL,
    commentable_id   BIGINT NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_comments_commentable ON comments (commentable_type, commentable_id);

CREATE TABLE attachments
(
    id              BIGSERIAL PRIMARY KEY,
    filename        TEXT   NOT NULL,
    filepath        TEXT   NOT NULL,
    filesize        BIGINT NOT NULL,
    filetype        TEXT   NOT NULL,
    user_id         BIGINT REFERENCES users (id) ON DELETE SET NULL,
    attachable_type TEXT   NOT NULL,
    attachable_id   BIGINT NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_attachments_attachable ON attachments (attachable_type, attachable_id);

CREATE TABLE notifications
(
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGINT REFERENCES users (id) ON DELETE CASCADE,
    title           TEXT    NOT NULL,
    content         TEXT    NOT NULL,
    type            TEXT    NOT NULL,
    read            BOOLEAN NOT NULL DEFAULT FALSE,
    notifiable_type TEXT    NOT NULL,
    notifiable_id   BIGINT  NOT NULL,
    created_at      TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    read_at         TIMESTAMP
);

CREATE INDEX idx_notifications_user ON notifications (user_id);
CREATE INDEX idx_notifications_notifiable ON notifications (notifiable_type, notifiable_id);


create table project_user
(
    user_id    BIGINT,
    project_id BIGINT
);


create table task_user
(
    user_id BIGINT,
    task_id BIGINT
);
