-- reverse: set comment to column: "finished_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."finished_at" IS '';
-- reverse: set comment to column: "started_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."started_at" IS '';
-- reverse: modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" DROP COLUMN "finished_at", DROP COLUMN "started_at";
