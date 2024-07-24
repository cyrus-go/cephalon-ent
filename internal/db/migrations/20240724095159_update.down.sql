-- reverse: set comment to column: "reason" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."reason" IS '';
-- reverse: modify "mission_failed_feedbacks" table
ALTER TABLE "mission_failed_feedbacks" DROP COLUMN "reason";
