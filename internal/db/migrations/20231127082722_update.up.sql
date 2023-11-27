-- modify "extra_service_prices" table
ALTER TABLE "extra_service_prices" ADD COLUMN "extra_service_type" character varying NOT NULL DEFAULT 'unknown';
-- set comment to column: "extra_service_type" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."extra_service_type" IS '附加服务类型';
