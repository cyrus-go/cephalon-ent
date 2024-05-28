-- reverse: set comment to column: "current_balance" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."current_balance" IS '';
-- reverse: modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" DROP COLUMN "current_balance";
