-- reverse: set comment to column: "status" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."status" IS '';
-- reverse: modify "mission_failed_feedbacks" table
ALTER TABLE "mission_failed_feedbacks" DROP COLUMN "status";
