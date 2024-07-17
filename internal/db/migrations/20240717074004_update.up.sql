-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "gpu_temperature" double precision NOT NULL DEFAULT 0, ADD COLUMN "cpu_temperature" double precision NOT NULL DEFAULT 0;
-- set comment to column: "gpu_temperature" on table: "devices"
COMMENT ON COLUMN "devices" ."gpu_temperature" IS 'GPU 温度(单位:℃)';
-- set comment to column: "cpu_temperature" on table: "devices"
COMMENT ON COLUMN "devices" ."cpu_temperature" IS 'CPU 温度(单位:℃)';
