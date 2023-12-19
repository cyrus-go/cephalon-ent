-- modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" ADD COLUMN "extra_service_billing_type" character varying NOT NULL DEFAULT 'unknown';
-- set comment to column: "extra_service_billing_type" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."extra_service_billing_type" IS '是否为计时类型任务';
