-- reverse: set comment to column: "cpus" on table: "devices"
COMMENT ON COLUMN "devices" ."cpus" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "cpus";
