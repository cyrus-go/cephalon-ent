-- reverse: set comment to column: "hmac_secret" on table: "users"
COMMENT ON COLUMN "users" ."hmac_secret" IS '';
-- reverse: set comment to column: "hmac_key" on table: "users"
COMMENT ON COLUMN "users" ."hmac_key" IS '';
-- reverse: set comment to column: "platform" on table: "users"
COMMENT ON COLUMN "users" ."platform" IS '';
-- reverse: set comment to column: "type" on table: "users"
COMMENT ON COLUMN "users" ."type" IS '';
-- reverse: set comment to column: "status" on table: "users"
COMMENT ON COLUMN "users" ."status" IS '';
-- reverse: set comment to column: "avatar_url" on table: "users"
COMMENT ON COLUMN "users" ."avatar_url" IS '';
-- reverse: set comment to column: "password" on table: "users"
COMMENT ON COLUMN "users" ."password" IS '';
-- reverse: set comment to column: "phone" on table: "users"
COMMENT ON COLUMN "users" ."phone" IS '';
-- reverse: set comment to column: "nick_name" on table: "users"
COMMENT ON COLUMN "users" ."nick_name" IS '';
-- reverse: set comment to column: "name" on table: "users"
COMMENT ON COLUMN "users" ."name" IS '';
-- reverse: set comment to table: "users"
COMMENT ON TABLE "users" IS '';
-- reverse: create index "user_hmac_key_deleted_at" to table: "users"
DROP INDEX "user_hmac_key_deleted_at";
-- reverse: create index "user_phone_deleted_at" to table: "users"
DROP INDEX "user_phone_deleted_at";
-- reverse: create index "user_name_deleted_at" to table: "users"
DROP INDEX "user_name_deleted_at";
-- reverse: create "users" table
DROP TABLE "users";
