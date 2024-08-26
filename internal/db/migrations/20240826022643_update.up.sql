-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "gift_mission_config_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "gift_mission_config_id" on table: "devices"
COMMENT ON COLUMN "devices" ."gift_mission_config_id" IS '外键补贴任务配置 id';
-- create "gift_mission_configs" table
CREATE TABLE "gift_mission_configs" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "stability_level" character varying NOT NULL DEFAULT 'good', "gpu_version" character varying NOT NULL DEFAULT 'RTX3080', "gap_base" bigint NOT NULL DEFAULT 0, "gap_random_max" bigint NOT NULL DEFAULT 0, "gap_random_min" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "giftmissionconfig_stability_level_gpu_version" to table: "gift_mission_configs"
CREATE INDEX "giftmissionconfig_stability_level_gpu_version" ON "gift_mission_configs" ("stability_level", "gpu_version");
-- set comment to table: "gift_mission_configs"
COMMENT ON TABLE "gift_mission_configs" IS '补贴任务参数配置（需要初始化数据）';
-- set comment to column: "id" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "stability_level" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."stability_level" IS '稳定性等级';
-- set comment to column: "gpu_version" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."gpu_version" IS 'GPU 版本';
-- set comment to column: "gap_base" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."gap_base" IS '间隔基数';
-- set comment to column: "gap_random_max" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."gap_random_max" IS '间隔随机范围最大值';
-- set comment to column: "gap_random_min" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."gap_random_min" IS '间隔随机范围最小值';
