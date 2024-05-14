-- modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" ADD COLUMN "cancel_reason" character varying NOT NULL DEFAULT '';
-- set comment to column: "cancel_reason" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."cancel_reason" IS '取消扣费原因';
