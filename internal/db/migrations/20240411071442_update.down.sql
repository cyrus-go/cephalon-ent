-- reverse: set comment to column: "operate_user_id" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."operate_user_id" IS '';
-- reverse: modify "transfer_orders" table
ALTER TABLE "transfer_orders" DROP COLUMN "operate_user_id";
