-- reverse: set comment to column: "gpu_version" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gpu_version" IS '';
-- reverse: modify "device_configs" table
ALTER TABLE "device_configs" DROP COLUMN "gpu_version";
