-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL, "nick_name" character varying NOT NULL DEFAULT '请设置昵称', "phone" character varying NOT NULL DEFAULT '', "password" character varying NOT NULL DEFAULT '', "avatar_url" character varying NOT NULL DEFAULT '', "status" character varying NOT NULL DEFAULT 'normal', "type" character varying NOT NULL DEFAULT 'personal', "platform" bigint NOT NULL DEFAULT 0, "hmac_key" character varying NOT NULL DEFAULT '', "hmac_secret" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- create index "user_name_deleted_at" to table: "users"
CREATE UNIQUE INDEX "user_name_deleted_at" ON "users" ("name", "deleted_at");
-- create index "user_phone_deleted_at" to table: "users"
CREATE UNIQUE INDEX "user_phone_deleted_at" ON "users" ("phone", "deleted_at");
-- create index "user_hmac_key_deleted_at" to table: "users"
CREATE UNIQUE INDEX "user_hmac_key_deleted_at" ON "users" ("hmac_key", "deleted_at");
-- set comment to table: "users"
COMMENT ON TABLE "users" IS '用户表，手机号唯一';
-- set comment to column: "name" on table: "users"
COMMENT ON COLUMN "users" ."name" IS '用户名';
-- set comment to column: "nick_name" on table: "users"
COMMENT ON COLUMN "users" ."nick_name" IS '用户昵称';
-- set comment to column: "phone" on table: "users"
COMMENT ON COLUMN "users" ."phone" IS '手机号';
-- set comment to column: "password" on table: "users"
COMMENT ON COLUMN "users" ."password" IS '密码';
-- set comment to column: "avatar_url" on table: "users"
COMMENT ON COLUMN "users" ."avatar_url" IS '头像路径';
-- set comment to column: "status" on table: "users"
COMMENT ON COLUMN "users" ."status" IS '用户状态';
-- set comment to column: "type" on table: "users"
COMMENT ON COLUMN "users" ."type" IS '用户类型';
-- set comment to column: "platform" on table: "users"
COMMENT ON COLUMN "users" ."platform" IS '用户可以在什么平台登录，二进制开关数据';
-- set comment to column: "hmac_key" on table: "users"
COMMENT ON COLUMN "users" ."hmac_key" IS '用户使用任务相关功能的密钥对的键，唯一';
-- set comment to column: "hmac_secret" on table: "users"
COMMENT ON COLUMN "users" ."hmac_secret" IS '用户使用任务相关功能的密钥对的值';
