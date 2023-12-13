-- modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" ADD COLUMN "mission_order_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "mission_order_id" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."mission_order_id" IS '任务订单 id，外键关联任务订单';
