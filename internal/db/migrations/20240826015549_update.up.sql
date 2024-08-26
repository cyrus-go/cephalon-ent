-- modify "device_configs" table
ALTER TABLE "device_configs" ADD COLUMN "gpu_version" character varying NOT NULL DEFAULT '';
-- set comment to column: "gpu_version" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gpu_version" IS 'GPU 版本';
