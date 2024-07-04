-- reverse: set comment to column: "user_status" on table: "users"
COMMENT ON COLUMN "users" ."user_status" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "user_status";
