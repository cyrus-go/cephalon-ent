-- modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" ADD COLUMN "deduct_standard" bigint NOT NULL DEFAULT 0;
-- set comment to column: "deduct_standard" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."deduct_standard" IS '扣费标准，单位：分';
