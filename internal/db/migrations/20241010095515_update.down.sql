-- reverse: set comment to column: "close_way" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."close_way" IS '';
-- reverse: modify "mission_load_balances" table
ALTER TABLE "mission_load_balances" DROP COLUMN "close_way";
