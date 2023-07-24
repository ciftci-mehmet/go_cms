CREATE TABLE "posts" (
                         "id" bigserial PRIMARY KEY,
                         "user_id" integer NOT NULL,
                         "title" varchar NOT NULL,
                         "body" text NOT NULL,
                         "status" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT(now()),
                         "updated_at" timestamptz
);

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX idx_posts_user_id ON "posts"("user_id");
CREATE INDEX idx_posts_title ON "posts"("title");
