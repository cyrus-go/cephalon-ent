-- modify "device_gpu_missions" table
ALTER TABLE "device_gpu_missions" ADD COLUMN "device_slot" smallint NOT NULL DEFAULT 0, ADD COLUMN "gpu_status" character varying NOT NULL DEFAULT 'offline';
-- set comment to column: "device_slot" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."device_slot" IS '显卡占用设备的插槽';
-- set comment to column: "gpu_status" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."gpu_status" IS 'gpu 当前状态';
