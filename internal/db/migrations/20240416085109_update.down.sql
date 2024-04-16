-- reverse: set comment to column: "target_symbol_id" on table: "bills"
COMMENT ON COLUMN "bills" ."target_symbol_id" IS '';
-- reverse: set comment to column: "symbol_id" on table: "bills"
COMMENT ON COLUMN "bills" ."symbol_id" IS '外键币种 id';
-- reverse: modify "bills" table
ALTER TABLE "bills" DROP COLUMN "target_symbol_id";
