-- modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" ADD COLUMN "reject_reason" character varying NOT NULL DEFAULT '';
-- set comment to column: "reject_reason" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."reject_reason" IS '拒绝扣费原因';
