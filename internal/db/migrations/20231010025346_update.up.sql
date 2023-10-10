-- modify "mission_orders" table
ALTER TABLE "mission_orders" ADD COLUMN "device_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "device_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."device_id" IS '关联的设备 id';
