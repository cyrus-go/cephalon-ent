-- reverse: set comment to column: "mission_tag" on table: "devices"
COMMENT ON COLUMN "devices" ."mission_tag" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "mission_tag";
