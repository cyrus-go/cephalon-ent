-- modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" ADD COLUMN "settled_amount" bigint NOT NULL DEFAULT 0, ADD COLUMN "settled_count" bigint NOT NULL DEFAULT 0, ADD COLUMN "total_settle_count" bigint NOT NULL DEFAULT 0, ADD COLUMN "lately_settled_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP;
-- set comment to column: "settled_amount" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."settled_amount" IS '已结算金额';
-- set comment to column: "settled_count" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."settled_count" IS '已结算次数';
-- set comment to column: "total_settle_count" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."total_settle_count" IS '总结算次数';
-- set comment to column: "lately_settled_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."lately_settled_at" IS '上一次结算时间';
