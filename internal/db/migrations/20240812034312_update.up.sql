-- modify "invoke_model_orders" table
ALTER TABLE "invoke_model_orders" ADD COLUMN "record_time" timestamptz NOT NULL;
-- set comment to column: "record_time" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."record_time" IS '记录时间';
