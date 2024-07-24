-- modify "mission_failed_feedbacks" table
ALTER TABLE "mission_failed_feedbacks" ADD COLUMN "reason" character varying NOT NULL DEFAULT '';
-- set comment to column: "reason" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."reason" IS '任务失败的原因';
