-- modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" ADD COLUMN "unit_cep" bigint NOT NULL DEFAULT 0;
-- set comment to column: "unit_cep" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."unit_cep" IS '任务单价，按次(count)就是 unit_cep/次，按时(time)就是 unit_cep/分钟';
