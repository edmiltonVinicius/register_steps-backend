CREATE TABLE IF NOT EXISTS users (
    "id" bigserial,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "first_name" text NOT NULL,
    "last_name" text NOT NULL,
    "email" text NOT NULL UNIQUE,
    "password" text NOT NULL,
    "country" text NOT NULL,
    PRIMARY KEY ("id")
);

COMMENT ON TABLE "users" IS 'Table of users';

COMMENT ON COLUMN "users"."first_name" IS 'First name the user';

COMMENT ON COLUMN "users"."last_name" IS 'Last name the user';

COMMENT ON COLUMN "users"."email" IS 'E-mail the user';

COMMENT ON COLUMN "users"."country" IS 'Country the user';

COMMENT ON COLUMN "users"."password" IS 'Hash password the user';