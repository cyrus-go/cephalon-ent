-- reverse: set comment to column: "withdraw_real_amount" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."withdraw_real_amount" IS '';
-- reverse: set comment to column: "withdraw_rate" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."withdraw_rate" IS '';
-- reverse: modify "transfer_orders" table
ALTER TABLE "transfer_orders" DROP COLUMN "withdraw_real_amount", DROP COLUMN "withdraw_rate";
