-- reverse: set comment to column: "type" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."type" IS '';
-- reverse: modify "mission_failed_feedbacks" table
ALTER TABLE "mission_failed_feedbacks" DROP COLUMN "type";
