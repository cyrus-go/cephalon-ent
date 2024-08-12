-- reverse: set comment to column: "record_time" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."record_time" IS '';
-- reverse: modify "invoke_model_orders" table
ALTER TABLE "invoke_model_orders" DROP COLUMN "record_time";
