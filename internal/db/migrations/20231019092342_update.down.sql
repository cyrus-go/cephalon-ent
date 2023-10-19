-- reverse: set comment to column: "mission_id" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."mission_id" IS '';
-- reverse: set comment to column: "max_online_mission" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."max_online_mission" IS '';
-- reverse: set comment to column: "able_mission_kind" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."able_mission_kind" IS '';
-- reverse: modify "device_gpu_missions" table
ALTER TABLE "device_gpu_missions" DROP COLUMN "mission_id", DROP COLUMN "max_online_mission", DROP COLUMN "able_mission_kind";
