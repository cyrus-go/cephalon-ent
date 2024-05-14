-- reverse: set comment to column: "reason" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."reason" IS '';
-- reverse: modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" DROP COLUMN "reason";
