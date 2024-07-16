-- set comment to table: "device_reboot_times"
COMMENT ON TABLE "device_reboot_times" IS '设备重启时间记录，已弃用';
-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "stability" bigint NOT NULL DEFAULT 0;
-- set comment to column: "stability" on table: "devices"
COMMENT ON COLUMN "devices" ."stability" IS '稳定性，数值越小越稳定';
-- create "device_states" table
CREATE TABLE "device_states" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "delay" double precision NOT NULL DEFAULT 0, "device_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "devicestate_device_id" to table: "device_states"
CREATE INDEX "devicestate_device_id" ON "device_states" ("device_id");
-- set comment to table: "device_states"
COMMENT ON TABLE "device_states" IS '设备状态信息 目前只有延迟信息';
-- set comment to column: "id" on table: "device_states"
COMMENT ON COLUMN "device_states" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "device_states"
COMMENT ON COLUMN "device_states" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "device_states"
COMMENT ON COLUMN "device_states" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "device_states"
COMMENT ON COLUMN "device_states" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "device_states"
COMMENT ON COLUMN "device_states" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "device_states"
COMMENT ON COLUMN "device_states" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "delay" on table: "device_states"
COMMENT ON COLUMN "device_states" ."delay" IS '延迟(单位:ms)';
-- set comment to column: "device_id" on table: "device_states"
COMMENT ON COLUMN "device_states" ."device_id" IS '外键设备 id';
