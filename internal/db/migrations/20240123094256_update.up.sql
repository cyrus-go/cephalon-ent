-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "warning_times" bigint NOT NULL DEFAULT 0;
-- set comment to column: "warning_times" on table: "missions"
COMMENT ON COLUMN "missions" ."warning_times" IS '预警次数，任务运行时间超过一定时间会发送预警消息';
