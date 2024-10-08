-- modify "mission_load_balances" table
ALTER TABLE "mission_load_balances" ADD COLUMN "current_mission_count" smallint NOT NULL DEFAULT 1;
-- set comment to column: "current_mission_count" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."current_mission_count" IS '当前应用数（调整中的以此为依据）';
