-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "manage_name" character varying NOT NULL DEFAULT '默认管理设备名称';
-- set comment to column: "manage_name" on table: "devices"
COMMENT ON COLUMN "devices" ."manage_name" IS '运维管理设备名称';
