-- reverse: set comment to column: "device_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."device_id" IS '';
-- reverse: modify "mission_orders" table
ALTER TABLE "mission_orders" DROP COLUMN "device_id";
