CREATE TABLE "agents" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "status" varchar NOT NULL DEFAULT 'INACTIVE',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tickets" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "status" varchar NOT NULL DEFAULT 'PENDING',
  "assigned_to" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "assignments" (
  "id" bigserial PRIMARY KEY,
  "ticket_id" bigint NOT NULL,
  "agent_id" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'PENDING',
  "assigned_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "agents" ("name");

CREATE INDEX ON "assignments" ("ticket_id");

CREATE INDEX ON "assignments" ("agent_id");

CREATE INDEX ON "assignments" ("ticket_id", "agent_id");

COMMENT ON COLUMN "agents"."status" IS 'Agent Status can be active or inactive';

COMMENT ON COLUMN "tickets"."status" IS 'Ticket Status can be PENDING or RESOLVED';

COMMENT ON COLUMN "assignments"."status" IS 'Ticket Status can be PENDING or RESOLVED';

ALTER TABLE "tickets" ADD FOREIGN KEY ("assigned_to") REFERENCES "agents" ("id");

ALTER TABLE "assignments" ADD FOREIGN KEY ("ticket_id") REFERENCES "tickets" ("id") ON DELETE CASCADE;

ALTER TABLE "assignments" ADD FOREIGN KEY ("agent_id") REFERENCES "agents" ("id") ON DELETE CASCADE;