-- reverse: set comment to column: "campaign_order_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."campaign_order_id" IS '';
-- reverse: create index "recharge_orders_campaign_order_id_key" to table: "recharge_orders"
DROP INDEX "recharge_orders_campaign_order_id_key";
-- reverse: modify "recharge_orders" table
ALTER TABLE "recharge_orders" DROP COLUMN "campaign_order_id";
