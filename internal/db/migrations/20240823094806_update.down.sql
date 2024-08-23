-- reverse: set comment to column: "device_id" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."device_id" IS '';
-- reverse: set comment to column: "gap_random_max" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gap_random_max" IS '';
-- reverse: set comment to column: "gap_base" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gap_base" IS '';
-- reverse: set comment to column: "deleted_at" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."id" IS '';
-- reverse: set comment to table: "device_configs"
COMMENT ON TABLE "device_configs" IS '';
-- reverse: create index "deviceconfig_device_id" to table: "device_configs"
DROP INDEX "deviceconfig_device_id";
-- reverse: create index "device_configs_device_id_key" to table: "device_configs"
DROP INDEX "device_configs_device_id_key";
-- reverse: create "device_configs" table
DROP TABLE "device_configs";
