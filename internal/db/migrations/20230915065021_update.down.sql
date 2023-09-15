-- reverse: set comment to column: "user_id" on table: "campaign_orders"
COMMENT ON COLUMN "campaign_orders" ."user_id" IS '';
-- reverse: set comment to column: "campaign_id" on table: "campaign_orders"
COMMENT ON COLUMN "campaign_orders" ."campaign_id" IS '';
-- reverse: create "campaign_orders" table
DROP TABLE "campaign_orders";
-- reverse: set comment to column: "campaign_order_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."campaign_order_id" IS '';
-- reverse: modify "cost_bills" table
ALTER TABLE "cost_bills" DROP COLUMN "campaign_order_id";
