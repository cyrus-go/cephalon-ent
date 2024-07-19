-- reverse: set comment to column: "old_mission_id" on table: "missions"
COMMENT ON COLUMN "missions" ."old_mission_id" IS '';
-- reverse: set comment to column: "use_auth" on table: "missions"
COMMENT ON COLUMN "missions" ."use_auth" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "old_mission_id", DROP COLUMN "use_auth";
