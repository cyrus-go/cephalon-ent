-- reverse: set comment to column: "google_id" on table: "users"
COMMENT ON COLUMN "users" ."google_id" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "google_id";
