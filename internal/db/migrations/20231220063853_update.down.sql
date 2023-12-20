-- reverse: set comment to column: "lately_settled_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."lately_settled_at" IS '';
-- reverse: set comment to column: "total_settle_count" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."total_settle_count" IS '';
-- reverse: set comment to column: "settled_count" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."settled_count" IS '';
-- reverse: set comment to column: "settled_amount" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."settled_amount" IS '';
-- reverse: modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" DROP COLUMN "lately_settled_at", DROP COLUMN "total_settle_count", DROP COLUMN "settled_count", DROP COLUMN "settled_amount";
