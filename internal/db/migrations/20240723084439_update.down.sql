-- reverse: set comment to column: "user_id" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."user_id" IS '';
-- reverse: set comment to column: "mission_id" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."mission_id" IS '';
-- reverse: set comment to column: "device_id" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."device_id" IS '';
-- reverse: set comment to column: "mission_name" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."mission_name" IS '';
-- reverse: set comment to column: "deleted_at" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."id" IS '';
-- reverse: set comment to table: "mission_failed_feedbacks"
COMMENT ON TABLE "mission_failed_feedbacks" IS '';
-- reverse: create index "missionfailedfeedback_mission_id" to table: "mission_failed_feedbacks"
DROP INDEX "missionfailedfeedback_mission_id";
-- reverse: create index "missionfailedfeedback_device_id" to table: "mission_failed_feedbacks"
DROP INDEX "missionfailedfeedback_device_id";
-- reverse: create index "missionfailedfeedback_user_id" to table: "mission_failed_feedbacks"
DROP INDEX "missionfailedfeedback_user_id";
-- reverse: create index "mission_failed_feedbacks_mission_id_key" to table: "mission_failed_feedbacks"
DROP INDEX "mission_failed_feedbacks_mission_id_key";
-- reverse: create "mission_failed_feedbacks" table
DROP TABLE "mission_failed_feedbacks";
