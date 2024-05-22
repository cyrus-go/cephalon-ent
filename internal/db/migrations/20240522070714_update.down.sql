-- reverse: set comment to column: "deduct_standard" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."deduct_standard" IS '';
-- reverse: modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" DROP COLUMN "deduct_standard";
