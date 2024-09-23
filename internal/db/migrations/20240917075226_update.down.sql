-- reverse: set comment to column: "github_id" on table: "users"
COMMENT ON COLUMN "users" ."github_id" IS '';
-- reverse: create index "user_github_id" to table: "users"
DROP INDEX "user_github_id";
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "github_id";
