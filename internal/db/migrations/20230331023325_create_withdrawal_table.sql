-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user_withdrawal" (
    "uid" UUID PRIMARY KEY NOT NULL DEFAULT "uuid_generate_v4" (),
    "from_user_uid" UUID NOT NULL,
    "to_user_uid" UUID NOT NULL,
    "money" DECIMAL NOT NULL,
    "status" TEXT NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamp,
    FOREIGN KEY (to_user_uid) REFERENCES "users" (uid),
    FOREIGN KEY (from_user_uid) REFERENCES "users" (uid)
);

-- +goose StatementEnd