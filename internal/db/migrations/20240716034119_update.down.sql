-- reverse: set comment to column: "device_id" on table: "device_states"
COMMENT ON COLUMN "device_states" ."device_id" IS '';
-- reverse: set comment to column: "delay" on table: "device_states"
COMMENT ON COLUMN "device_states" ."delay" IS '';
-- reverse: set comment to column: "deleted_at" on table: "device_states"
COMMENT ON COLUMN "device_states" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "device_states"
COMMENT ON COLUMN "device_states" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "device_states"
COMMENT ON COLUMN "device_states" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "device_states"
COMMENT ON COLUMN "device_states" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "device_states"
COMMENT ON COLUMN "device_states" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "device_states"
COMMENT ON COLUMN "device_states" ."id" IS '';
-- reverse: set comment to table: "device_states"
COMMENT ON TABLE "device_states" IS '';
-- reverse: create index "devicestate_device_id" to table: "device_states"
DROP INDEX "devicestate_device_id";
-- reverse: create "device_states" table
DROP TABLE "device_states";
-- reverse: set comment to column: "stability" on table: "devices"
COMMENT ON COLUMN "devices" ."stability" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "stability";
-- reverse: set comment to table: "device_reboot_times"
COMMENT ON TABLE "device_reboot_times" IS '设备重启时间记录';
