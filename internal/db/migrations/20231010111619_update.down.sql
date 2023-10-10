-- reverse: set comment to column: "pop_version" on table: "users"
COMMENT ON COLUMN "users" ."pop_version" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "pop_version";
