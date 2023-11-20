-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "name" character varying NOT NULL DEFAULT '', ADD COLUMN "type" character varying NOT NULL DEFAULT 'ordinary', ADD COLUMN "cores_number" bigint NOT NULL DEFAULT 0, ADD COLUMN "cpu" character varying NOT NULL DEFAULT '', ADD COLUMN "memory" bigint NOT NULL DEFAULT 0, ADD COLUMN "disk" bigint NOT NULL DEFAULT 0;
-- set comment to column: "name" on table: "devices"
COMMENT ON COLUMN "devices" ."name" IS '设备名称';
-- set comment to column: "type" on table: "devices"
COMMENT ON COLUMN "devices" ."type" IS '设备类型';
-- set comment to column: "cores_number" on table: "devices"
COMMENT ON COLUMN "devices" ."cores_number" IS '核心数';
-- set comment to column: "cpu" on table: "devices"
COMMENT ON COLUMN "devices" ."cpu" IS 'CPU型号';
-- set comment to column: "memory" on table: "devices"
COMMENT ON COLUMN "devices" ."memory" IS '内存(单位:G)';
-- set comment to column: "disk" on table: "devices"
COMMENT ON COLUMN "devices" ."disk" IS '硬盘(单位:T)';
-- modify "mission_productions" table
ALTER TABLE "mission_productions" ADD COLUMN "device_slot" smallint NOT NULL DEFAULT 0;
-- set comment to column: "device_slot" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."device_slot" IS '显卡占用设备的插槽';
