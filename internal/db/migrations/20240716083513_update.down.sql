-- reverse: set comment to column: "device_id" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."device_id" IS '';
-- reverse: set comment to column: "deleted_at" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."id" IS '';
-- reverse: set comment to table: "device_offline_records"
COMMENT ON TABLE "device_offline_records" IS '';
-- reverse: create index "deviceofflinerecord_device_id" to table: "device_offline_records"
DROP INDEX "deviceofflinerecord_device_id";
-- reverse: create "device_offline_records" table
DROP TABLE "device_offline_records";
