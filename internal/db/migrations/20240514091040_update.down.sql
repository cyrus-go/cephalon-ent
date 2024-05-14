-- reverse: set comment to column: "cancel_reason" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."cancel_reason" IS '';
-- reverse: modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" DROP COLUMN "cancel_reason";
