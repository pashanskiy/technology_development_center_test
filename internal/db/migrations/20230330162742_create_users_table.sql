-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "clients" (
    "uid" UUID PRIMARY KEY NOT NULL DEFAULT "uuid_generate_v4" (),
    "money" DECIMAL NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp
);

-- +goose StatementEnd