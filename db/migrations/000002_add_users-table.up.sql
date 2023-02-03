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
ALTER TABLE "agents" ADD COLUMN "user_id" SERIAL NOT NULL;

ALTER TABLE "agents" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "agents" ADD CONSTRAINT "name_user_id_key" UNIQUE ("name", "user_id");

-- add createdby column to tickets
ALTER TABLE "tickets" ADD COLUMN "createdby_id" SERIAL NOT NULL;

ALTER TABLE "tickets" ADD FOREIGN KEY ("createdby_id") REFERENCES "users" ("id");

ALTER TABLE "tickets" ADD CONSTRAINT "createdby_id_assignedto_id_key" UNIQUE ("assigned_to", "createdby_id");
