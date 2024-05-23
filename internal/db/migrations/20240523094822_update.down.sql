-- reverse: set comment to column: "bound_at" on table: "users"
COMMENT ON COLUMN "users" ."bound_at" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "bound_at";
