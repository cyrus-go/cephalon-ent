-- reverse: set comment to column: "now_time" on table: "device_reboot_times"
COMMENT ON COLUMN "device_reboot_times" ."now_time" IS '';
-- reverse: modify "device_reboot_times" table
ALTER TABLE "device_reboot_times" DROP COLUMN "now_time";
