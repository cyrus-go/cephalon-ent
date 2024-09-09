-- reverse: set comment to column: "gpu_num" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."gpu_num" IS '';
-- reverse: set comment to column: "device_slots" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."device_slots" IS '';
-- reverse: modify "mission_productions" table
ALTER TABLE "mission_productions" DROP COLUMN "gpu_num", DROP COLUMN "device_slots";
