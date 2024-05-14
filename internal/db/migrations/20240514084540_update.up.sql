-- modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" ADD COLUMN "reason" character varying NOT NULL DEFAULT '';
-- set comment to column: "reason" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."reason" IS '扣费原因';
