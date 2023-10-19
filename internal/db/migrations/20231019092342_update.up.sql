-- modify "device_gpu_missions" table
ALTER TABLE "device_gpu_missions" ADD COLUMN "able_mission_kind" bytea NULL, ADD COLUMN "max_online_mission" smallint NOT NULL DEFAULT 0, ADD COLUMN "mission_id" bytea NULL;
-- set comment to column: "able_mission_kind" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."able_mission_kind" IS '可以接的任务类型';
-- set comment to column: "max_online_mission" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."max_online_mission" IS '最大同时在线任务';
-- set comment to column: "mission_id" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."mission_id" IS '正在做的任务 id';
