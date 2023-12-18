-- reverse: set comment to column: "unit_cep" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."unit_cep" IS '';
-- reverse: modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" DROP COLUMN "unit_cep";
