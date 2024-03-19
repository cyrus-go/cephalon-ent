-- reverse: set comment to column: "withdraw_account" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."withdraw_account" IS '';
-- reverse: modify "transfer_orders" table
ALTER TABLE "transfer_orders" DROP COLUMN "withdraw_account";
