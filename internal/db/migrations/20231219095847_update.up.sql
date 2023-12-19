-- modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" ADD COLUMN "started_at" timestamptz NULL, ADD COLUMN "finished_at" timestamptz NULL;
-- set comment to column: "started_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."started_at" IS '附加服务开始执行时刻';
-- set comment to column: "finished_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."finished_at" IS '附加服务结束执行时刻';
