-- create "device_offline_records" table
CREATE TABLE "device_offline_records" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "device_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "deviceofflinerecord_device_id" to table: "device_offline_records"
CREATE INDEX "deviceofflinerecord_device_id" ON "device_offline_records" ("device_id");
-- set comment to table: "device_offline_records"
COMMENT ON TABLE "device_offline_records" IS '设备离线时间记录';
-- set comment to column: "id" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "device_id" on table: "device_offline_records"
COMMENT ON COLUMN "device_offline_records" ."device_id" IS '外键设备 id';
