-- reverse: set comment to column: "cloud_space" on table: "users"
COMMENT ON COLUMN "users" ."cloud_space" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "cloud_space";
