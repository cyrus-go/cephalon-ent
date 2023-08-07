-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "name" character varying NOT NULL DEFAULT '请设置昵称', "phone" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- create index "user_phone_deleted_at" to table: "users"
CREATE UNIQUE INDEX "user_phone_deleted_at" ON "users" ("phone", "deleted_at");
-- set comment to column: "name" on table: "users"
COMMENT ON COLUMN "users" ."name" IS '用户名';
-- set comment to column: "phone" on table: "users"
COMMENT ON COLUMN "users" ."phone" IS '手机号';
