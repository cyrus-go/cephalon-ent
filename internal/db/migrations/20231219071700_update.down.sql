-- reverse: set comment to column: "extra_service_billing_type" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."extra_service_billing_type" IS '';
-- reverse: modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" DROP COLUMN "extra_service_billing_type";
