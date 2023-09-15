-- modify "cost_bills" table
ALTER TABLE "cost_bills" ADD COLUMN "campaign_order_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "campaign_order_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."campaign_order_id" IS '活动订单 id';
-- create "campaign_orders" table
CREATE TABLE "campaign_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "campaign_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "campaign_id" on table: "campaign_orders"
COMMENT ON COLUMN "campaign_orders" ."campaign_id" IS '活动 id';
-- set comment to column: "user_id" on table: "campaign_orders"
COMMENT ON COLUMN "campaign_orders" ."user_id" IS '用户 id';
