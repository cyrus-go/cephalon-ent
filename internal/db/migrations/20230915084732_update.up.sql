-- modify "recharge_orders" table
ALTER TABLE "recharge_orders" ADD COLUMN "campaign_order_id" bigint NULL;
-- create index "recharge_orders_campaign_order_id_key" to table: "recharge_orders"
CREATE UNIQUE INDEX "recharge_orders_campaign_order_id_key" ON "recharge_orders" ("campaign_order_id");
-- set comment to column: "campaign_order_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."campaign_order_id" IS '活动订单 id';
