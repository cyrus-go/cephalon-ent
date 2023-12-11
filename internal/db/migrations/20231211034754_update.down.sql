-- reverse: set comment to column: "mission_order_id" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."mission_order_id" IS '';
-- reverse: modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" DROP COLUMN "mission_order_id";
