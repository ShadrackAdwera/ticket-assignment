-- CREATE users table
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- add userId column to agents
ALTER TABLE "agents" ADD COLUMN "user_id" SERIAL;

-- make column unique index and foreign key
ALTER TABLE "agents" ADD CONSTRAINT "name_user_id_key" FOREIGN KEY ("user_id") REFERENCES "users" ("id");