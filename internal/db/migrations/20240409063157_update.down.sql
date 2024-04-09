-- reverse: set comment to column: "device_id" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."device_id" IS '';
-- reverse: set comment to column: "offline_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."offline_time" IS '';
-- reverse: set comment to column: "online_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."online_time" IS '';
-- reverse: set comment to column: "end_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."end_time" IS '';
-- reverse: set comment to column: "start_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."start_time" IS '';
-- reverse: set comment to column: "deleted_at" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."id" IS '';
-- reverse: set comment to table: "device_reboot_times"
COMMENT ON TABLE "device_reboot_times" IS '';
-- reverse: create index "devicereboottime_device_id" to table: "device_reboot_times"
DROP INDEX "devicereboottime_device_id";
-- reverse: create "device_reboot_times" table
DROP TABLE "device_reboot_times";
