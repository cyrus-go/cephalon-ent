-- reverse: set comment to column: "produce_user_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."produce_user_id" IS '';
-- reverse: set comment to column: "consume_user_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."consume_user_id" IS '';
-- reverse: set comment to column: "symbol_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."symbol_id" IS '';
-- reverse: set comment to column: "mission_batch_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_batch_id" IS '';
-- reverse: set comment to column: "mission_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_id" IS '';
-- reverse: set comment to column: "mission_batch_number" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_batch_number" IS '';
-- reverse: set comment to column: "finished_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."finished_at" IS '';
-- reverse: set comment to column: "started_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."started_at" IS '';
-- reverse: set comment to column: "serial_number" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."serial_number" IS '';
-- reverse: set comment to column: "call_way" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."call_way" IS '';
-- reverse: set comment to column: "mission_billing_type" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_billing_type" IS '';
-- reverse: set comment to column: "mission_type" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_type" IS '';
-- reverse: set comment to column: "gas_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."gas_amount" IS '';
-- reverse: set comment to column: "produce_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."produce_amount" IS '';
-- reverse: set comment to column: "consume_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."consume_amount" IS '';
-- reverse: set comment to column: "status" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."status" IS '';
-- reverse: set comment to table: "mission_orders"
COMMENT ON TABLE "mission_orders" IS '';
-- reverse: create "mission_orders" table
DROP TABLE "mission_orders";
-- reverse: modify "bills" table
ALTER TABLE "bills" DROP COLUMN "mission_order_bills";
