-- modify "mission_failed_feedbacks" table
ALTER TABLE "mission_failed_feedbacks" ADD COLUMN "type" character varying NOT NULL DEFAULT 'unknown';
-- set comment to column: "type" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."type" IS '任务失败反馈类型';
