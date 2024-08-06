-- reverse: set comment to column: "fault" on table: "devices"
COMMENT ON COLUMN "devices" ."fault" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "fault";
