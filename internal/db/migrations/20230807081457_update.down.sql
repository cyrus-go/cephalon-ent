-- reverse: set comment to column: "phone" on table: "users"
COMMENT ON COLUMN "users" ."phone" IS '';
-- reverse: set comment to column: "name" on table: "users"
COMMENT ON COLUMN "users" ."name" IS '';
-- reverse: create index "user_phone_deleted_at" to table: "users"
DROP INDEX "user_phone_deleted_at";
-- reverse: create "users" table
DROP TABLE "users";
