CREATE TABLE "users" (
  "id" serial UNIQUE PRIMARY KEY,
  "email" text UNIQUE NOT NULL,
  "username" text NOT NULL,
  "hash" bytea NOT NULL
);

CREATE TABLE "rfcs" (
  "id" serial UNIQUE PRIMARY KEY,
  "author_id" integer REFERENCES "users" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  "message" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "archived" boolean NOT NULL DEFAULT FALSE
);

CREATE TABLE "comments" (
  "id" serial UNIQUE PRIMARY KEY,
  "author_id" integer REFERENCES "users" ("id") ON DELETE SET NULL ON UPDATE CASCADE,
  "rfc_id" integer REFERENCES "rfcs" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "rfcs" ("id");

CREATE INDEX ON "rfcs" ("author_id");
