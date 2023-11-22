-- reverse: modify "mission_orders" table
ALTER TABLE "mission_orders" ALTER COLUMN "last_settled_at" SET DEFAULT '0001-01-01 08:05:43+08:05:43';
-- reverse: set comment to column: "profit_symbol_id" on table: "bills"
COMMENT ON COLUMN "bills" ."profit_symbol_id" IS '';
-- reverse: modify "bills" table
ALTER TABLE "bills" DROP COLUMN "profit_symbol_id";
