-- reverse: set comment to column: "nick_name" on table: "users"
COMMENT ON COLUMN "users" ."nick_name" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "nick_name";
