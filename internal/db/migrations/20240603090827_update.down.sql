-- reverse: set comment to column: "manage_name" on table: "devices"
COMMENT ON COLUMN "devices" ."manage_name" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "manage_name";
