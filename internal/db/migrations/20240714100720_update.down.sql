-- reverse: set comment to column: "applet_parent_id" on table: "users"
COMMENT ON COLUMN "users" ."applet_parent_id" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "applet_parent_id";
