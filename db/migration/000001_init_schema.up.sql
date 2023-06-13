CREATE TABLE "user"(
    "id" bigserial PRIMARY KEY,
    "username" varchar NOT NULL UNIQUE,
    "hashed_password" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT(now())
);

CREATE INDEX ON "user"("username");
