-- create "mission_failed_feedbacks" table
CREATE TABLE "mission_failed_feedbacks" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "mission_name" character varying NOT NULL DEFAULT '', "device_id" bigint NOT NULL DEFAULT 0, "mission_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "mission_failed_feedbacks_mission_id_key" to table: "mission_failed_feedbacks"
CREATE UNIQUE INDEX "mission_failed_feedbacks_mission_id_key" ON "mission_failed_feedbacks" ("mission_id");
-- create index "missionfailedfeedback_user_id" to table: "mission_failed_feedbacks"
CREATE INDEX "missionfailedfeedback_user_id" ON "mission_failed_feedbacks" ("user_id");
-- create index "missionfailedfeedback_device_id" to table: "mission_failed_feedbacks"
CREATE INDEX "missionfailedfeedback_device_id" ON "mission_failed_feedbacks" ("device_id");
-- create index "missionfailedfeedback_mission_id" to table: "mission_failed_feedbacks"
CREATE INDEX "missionfailedfeedback_mission_id" ON "mission_failed_feedbacks" ("mission_id");
-- set comment to table: "mission_failed_feedbacks"
COMMENT ON TABLE "mission_failed_feedbacks" IS '应用启动失败反馈信息表';
-- set comment to column: "id" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "mission_name" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."mission_name" IS '应用名称';
-- set comment to column: "device_id" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."device_id" IS '外键，反馈关联的设备 ID';
-- set comment to column: "mission_id" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."mission_id" IS '外键，反馈关联的任务 ID';
-- set comment to column: "user_id" on table: "mission_failed_feedbacks"
COMMENT ON COLUMN "mission_failed_feedbacks" ."user_id" IS '外键，反馈的用户 ID';
