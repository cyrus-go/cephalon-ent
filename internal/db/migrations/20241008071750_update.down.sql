-- reverse: set comment to column: "state" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."state" IS '';
-- reverse: modify "mission_load_balances" table
ALTER TABLE "mission_load_balances" DROP COLUMN "state";
