-- reverse: set comment to column: "device_slot" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."device_slot" IS '';
-- reverse: modify "mission_productions" table
ALTER TABLE "mission_productions" DROP COLUMN "device_slot";
-- reverse: set comment to column: "disk" on table: "devices"
COMMENT ON COLUMN "devices" ."disk" IS '';
-- reverse: set comment to column: "memory" on table: "devices"
COMMENT ON COLUMN "devices" ."memory" IS '';
-- reverse: set comment to column: "cpu" on table: "devices"
COMMENT ON COLUMN "devices" ."cpu" IS '';
-- reverse: set comment to column: "cores_number" on table: "devices"
COMMENT ON COLUMN "devices" ."cores_number" IS '';
-- reverse: set comment to column: "type" on table: "devices"
COMMENT ON COLUMN "devices" ."type" IS '';
-- reverse: set comment to column: "name" on table: "devices"
COMMENT ON COLUMN "devices" ."name" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "disk", DROP COLUMN "memory", DROP COLUMN "cpu", DROP COLUMN "cores_number", DROP COLUMN "type", DROP COLUMN "name";
