-- reverse: set comment to column: "gift_type" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."gift_type" IS '';
-- reverse: modify "transfer_orders" table
ALTER TABLE "transfer_orders" DROP COLUMN "gift_type";
