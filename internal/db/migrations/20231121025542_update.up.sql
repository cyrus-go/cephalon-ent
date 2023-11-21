-- modify "mission_orders" table
ALTER TABLE "mission_orders" ADD COLUMN "total_amount" bigint NOT NULL DEFAULT 0, ADD COLUMN "settled_amount" bigint NOT NULL DEFAULT 0, ADD COLUMN "settled_count" bigint NOT NULL DEFAULT 0, ADD COLUMN "total_settle_count" bigint NOT NULL DEFAULT 0, ADD COLUMN "last_settled_at" timestamptz default '0001-01-01 00:00:00.0000000 +00:00' NOT NULL;
-- set comment to column: "total_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."total_amount" IS '订单总金额';
-- set comment to column: "settled_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."settled_amount" IS '已结算金额';
-- set comment to column: "settled_count" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."settled_count" IS '已结算次数';
-- set comment to column: "total_settle_count" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."total_settle_count" IS '总结算次数';
-- set comment to column: "last_settled_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."last_settled_at" IS '上一次结算时间';
