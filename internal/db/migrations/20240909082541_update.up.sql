-- modify "mission_productions" table
ALTER TABLE "mission_productions" ADD COLUMN "device_slots" character varying NOT NULL DEFAULT '', ADD COLUMN "gpu_num" smallint NOT NULL DEFAULT 0;
-- set comment to column: "device_slots" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."device_slots" IS '显卡占用设备的插槽(多gpu)';
-- set comment to column: "gpu_num" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."gpu_num" IS '使用显卡数量';
