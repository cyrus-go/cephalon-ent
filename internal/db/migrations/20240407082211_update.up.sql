-- create "device_reboot_times" table
CREATE TABLE "device_reboot_times" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "start_time" timestamptz NOT NULL, "end_time" timestamptz NOT NULL, "online_time" character varying NOT NULL DEFAULT 'None', "offline_time" character varying NOT NULL DEFAULT 'None', "device_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "devicereboottime_device_id" to table: "device_reboot_times"
CREATE INDEX "devicereboottime_device_id" ON "device_reboot_times" ("device_id");
-- set comment to table: "device_reboot_times"
COMMENT ON TABLE "device_reboot_times" IS '设备重启时间记录';
-- set comment to column: "id" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "start_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."start_time" IS '设备开机时间';
-- set comment to column: "end_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."end_time" IS '设备关机时间';
-- set comment to column: "online_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."online_time" IS '设备运行时间';
-- set comment to column: "offline_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."offline_time" IS '设备宕机时间';
-- set comment to column: "device_id" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."device_id" IS '外键设备 id';
