-- reverse: set comment to column: "email" on table: "users"
COMMENT ON COLUMN "users" ."email" IS '';
-- reverse: set comment to column: "area_code" on table: "users"
COMMENT ON COLUMN "users" ."area_code" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "email", DROP COLUMN "area_code";
