-- Create "users" table
CREATE TABLE "public"."users" (
 "id" bigserial NOT NULL,
 "created_at" timestamptz NULL,
 "updated_at" timestamptz NULL,
 "deleted_at" timestamptz NULL,
 "email" text NOT NULL,
 "password" text NOT NULL,
 "full_name" character varying(255) NULL DEFAULT '',
 PRIMARY KEY ("id"),
 CONSTRAINT "uni_users_email" UNIQUE ("email")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create "refresh_token_families" table
CREATE TABLE "public"."refresh_token_families" (
 "id" bigserial NOT NULL,
 "created_at" timestamptz NULL,
 "updated_at" timestamptz NULL,
 "deleted_at" timestamptz NULL,
 "status" text NULL DEFAULT 'active',
 "user_id" bigint NOT NULL,
 PRIMARY KEY ("id"),
 CONSTRAINT "fk_users_refresh_tokens" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_refresh_token_families_deleted_at" to table: "refresh_token_families"
CREATE INDEX "idx_refresh_token_families_deleted_at" ON "public"."refresh_token_families" ("deleted_at");
-- Create "refresh_tokens" table
CREATE TABLE "public"."refresh_tokens" (
 "id" bigserial NOT NULL,
 "created_at" timestamptz NULL,
 "updated_at" timestamptz NULL,
 "deleted_at" timestamptz NULL,
 "jti" text NOT NULL,
 "status" text NULL DEFAULT 'inuse',
 "parent_id" bigint NULL,
 "refresh_token_family_id" bigint NOT NULL,
 PRIMARY KEY ("id"),
 CONSTRAINT "uni_refresh_tokens_jti" UNIQUE ("jti"),
 CONSTRAINT "fk_refresh_token_families_refresh_tokens" FOREIGN KEY ("refresh_token_family_id") REFERENCES "public"."refresh_token_families" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
 CONSTRAINT "fk_refresh_tokens_parent" FOREIGN KEY ("parent_id") REFERENCES "public"."refresh_tokens" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_refresh_tokens_deleted_at" to table: "refresh_tokens"
CREATE INDEX "idx_refresh_tokens_deleted_at" ON "public"."refresh_tokens" ("deleted_at");
