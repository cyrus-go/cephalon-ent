-- reverse: set comment to column: "current_mission_count" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."current_mission_count" IS '';
-- reverse: modify "mission_load_balances" table
ALTER TABLE "mission_load_balances" DROP COLUMN "current_mission_count";
