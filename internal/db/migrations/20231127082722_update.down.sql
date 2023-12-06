-- reverse: set comment to column: "extra_service_type" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."extra_service_type" IS '';
-- reverse: modify "extra_service_prices" table
ALTER TABLE "extra_service_prices" DROP COLUMN "extra_service_type";
