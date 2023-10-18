-- reverse: set comment to column: "gpu_status" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."gpu_status" IS '';
-- reverse: set comment to column: "device_slot" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."device_slot" IS '';
-- reverse: modify "device_gpu_missions" table
ALTER TABLE "device_gpu_missions" DROP COLUMN "gpu_status", DROP COLUMN "device_slot";
