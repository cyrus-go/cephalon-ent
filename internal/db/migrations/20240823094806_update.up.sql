-- create "device_configs" table
CREATE TABLE "device_configs" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "gap_base" bigint NOT NULL DEFAULT 0, "gap_random_max" bigint NOT NULL DEFAULT 0, "device_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "device_configs_device_id_key" to table: "device_configs"
CREATE UNIQUE INDEX "device_configs_device_id_key" ON "device_configs" ("device_id");
-- create index "deviceconfig_device_id" to table: "device_configs"
CREATE INDEX "deviceconfig_device_id" ON "device_configs" ("device_id");
-- set comment to table: "device_configs"
COMMENT ON TABLE "device_configs" IS '设备配置信息';
-- set comment to column: "id" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "gap_base" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gap_base" IS '间隔基数';
-- set comment to column: "gap_random_max" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gap_random_max" IS '间隔随机范围';
-- set comment to column: "device_id" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."device_id" IS '外键设备 id';
