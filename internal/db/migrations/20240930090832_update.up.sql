-- create "mission_load_balances" table
CREATE TABLE "mission_load_balances" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "mission_type" character varying NOT NULL DEFAULT 'unknown', "user_id" bigint NOT NULL, "started_at" timestamptz NOT NULL, "finished_at" timestamptz NOT NULL, "gpu_version" character varying NOT NULL DEFAULT 'unknown', "gpu_num" smallint NOT NULL DEFAULT 0, "max_mission_count" smallint NOT NULL DEFAULT 1, "min_mission_count" smallint NOT NULL DEFAULT 1, "mission_batch_id" bigint NOT NULL DEFAULT 0, "mission_batch_number" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- create index "missionloadbalance_mission_batch_id" to table: "mission_load_balances"
CREATE INDEX "missionloadbalance_mission_batch_id" ON "mission_load_balances" ("mission_batch_id");
-- create index "missionloadbalance_mission_batch_number" to table: "mission_load_balances"
CREATE INDEX "missionloadbalance_mission_batch_number" ON "mission_load_balances" ("mission_batch_number");
-- set comment to table: "mission_load_balances"
COMMENT ON TABLE "mission_load_balances" IS '任务负载均衡，多个应用自动启动，访问多的时候多启点，访问少的时候减少点';
-- set comment to column: "id" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "mission_type" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."mission_type" IS '任务类型';
-- set comment to column: "user_id" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."user_id" IS '用户 ID';
-- set comment to column: "started_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."started_at" IS '任务开始时刻';
-- set comment to column: "finished_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."finished_at" IS '任务完成时刻';
-- set comment to column: "gpu_version" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."gpu_version" IS '任务使用什么显卡在执行';
-- set comment to column: "gpu_num" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."gpu_num" IS '使用显卡数量';
-- set comment to column: "max_mission_count" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."max_mission_count" IS '应用数浮动上限';
-- set comment to column: "min_mission_count" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."min_mission_count" IS '应用数浮动下限';
-- set comment to column: "mission_batch_id" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."mission_batch_id" IS '外键关联任务批次';
-- set comment to column: "mission_batch_number" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."mission_batch_number" IS '任务批次号';
-- create "mission_load_balance_accesses" table
CREATE TABLE "mission_load_balance_accesses" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "mission_id" bigint NOT NULL, "mission_load_balance_id" bigint NOT NULL, "last_access" timestamptz NOT NULL, "access_count" integer NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "missionloadbalanceaccess_mission_id" to table: "mission_load_balance_accesses"
CREATE INDEX "missionloadbalanceaccess_mission_id" ON "mission_load_balance_accesses" ("mission_id");
-- create index "missionloadbalanceaccess_mission_load_balance_id" to table: "mission_load_balance_accesses"
CREATE INDEX "missionloadbalanceaccess_mission_load_balance_id" ON "mission_load_balance_accesses" ("mission_load_balance_id");
-- set comment to table: "mission_load_balance_accesses"
COMMENT ON TABLE "mission_load_balance_accesses" IS 'mission负载均衡访问记录';
-- set comment to column: "id" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "mission_id" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."mission_id" IS '任务 ID';
-- set comment to column: "mission_load_balance_id" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."mission_load_balance_id" IS '父 loadbanlace ID';
-- set comment to column: "last_access" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."last_access" IS '上次获取 api 时间';
-- set comment to column: "access_count" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."access_count" IS '已访问次数';
