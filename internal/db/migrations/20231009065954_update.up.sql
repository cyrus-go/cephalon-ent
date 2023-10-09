-- modify "mission_orders" table
ALTER TABLE "mission_orders" ADD COLUMN "buy_duration" bigint NOT NULL DEFAULT 0;
-- set comment to column: "buy_duration" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."buy_duration" IS '包时任务订单购买的时长（单位：小时）';
