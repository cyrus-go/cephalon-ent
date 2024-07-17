-- reverse: set comment to column: "cpu_temperature" on table: "devices"
COMMENT ON COLUMN "devices" ."cpu_temperature" IS '';
-- reverse: set comment to column: "gpu_temperature" on table: "devices"
COMMENT ON COLUMN "devices" ."gpu_temperature" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "cpu_temperature", DROP COLUMN "gpu_temperature";
