-- reverse: set comment to column: "symbol_id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."symbol_id" IS '';
-- reverse: modify "income_manages" table
ALTER TABLE "income_manages" DROP COLUMN "symbol_id";
