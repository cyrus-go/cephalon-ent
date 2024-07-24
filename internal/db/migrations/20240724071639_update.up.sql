-- modify "mission_failed_feedbacks" table
ALTER TABLE "mission_failed_feedbacks" ADD COLUMN "status" character varying NOT NULL DEFAULT 'init';
-- set comment to column: "status" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."status" IS '状态';
