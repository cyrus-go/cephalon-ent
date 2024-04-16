-- modify "device_reboot_times" table
ALTER TABLE "device_reboot_times" ADD COLUMN "now_time" timestamptz NOT NULL;
-- set comment to column: "now_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."now_time" IS '设备上线时间';
