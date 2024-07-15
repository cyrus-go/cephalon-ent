-- reverse: set comment to column: "temperature" on table: "devices"
COMMENT ON COLUMN "devices" ."temperature" IS '';
-- reverse: set comment to column: "delay" on table: "devices"
COMMENT ON COLUMN "devices" ."delay" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "temperature", DROP COLUMN "delay";
