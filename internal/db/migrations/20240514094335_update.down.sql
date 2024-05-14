-- reverse: set comment to column: "reject_reason" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."reject_reason" IS '';
-- reverse: modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" DROP COLUMN "reject_reason";
