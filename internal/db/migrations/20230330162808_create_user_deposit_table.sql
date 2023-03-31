-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user_deposit" (
  "uid" UUID PRIMARY KEY NOT NULL DEFAULT "uuid_generate_v4" (),
  "to_user_uid" UUID NOT NULL,
  "money" DECIMAL NOT NULL,
  "status" TEXT NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp,
  FOREIGN KEY (to_user_uid) REFERENCES "users" (uid)
);

-- +goose StatementEnd