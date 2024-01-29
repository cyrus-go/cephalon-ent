-- create "lottos" table
CREATE TABLE "lottos" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', "total_weight" bigint NOT NULL DEFAULT 0, "started_at" timestamptz NOT NULL, "ended_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'unknown', PRIMARY KEY ("id"));
-- create index "lotto_name" to table: "lottos"
CREATE UNIQUE INDEX "lotto_name" ON "lottos" ("name");
-- set comment to table: "lottos"
COMMENT ON TABLE "lottos" IS '抽奖活动表';
-- set comment to column: "id" on table: "lottos"
COMMENT ON COLUMN "lottos" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "lottos"
COMMENT ON COLUMN "lottos" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "lottos"
COMMENT ON COLUMN "lottos" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "name" on table: "lottos"
COMMENT ON COLUMN "lottos" ."name" IS '抽奖活动名称';
-- set comment to column: "total_weight" on table: "lottos"
COMMENT ON COLUMN "lottos" ."total_weight" IS '活动总权重';
-- set comment to column: "started_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."started_at" IS '活动开始时间';
-- set comment to column: "ended_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."ended_at" IS '活动结束时间';
-- set comment to column: "status" on table: "lottos"
COMMENT ON COLUMN "lottos" ."status" IS '状态';
-- create "lotto_get_count_records" table
CREATE TABLE "lotto_get_count_records" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "count" bigint NOT NULL DEFAULT 0, "type" character varying NOT NULL DEFAULT 'unknow', "recharge_amount" bigint NOT NULL DEFAULT 0, "lotto_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "lotto_get_count_records"
COMMENT ON TABLE "lotto_get_count_records" IS '用户获得抽奖次数记录表';
-- set comment to column: "id" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "count" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."count" IS '此次奖励的抽奖次数';
-- set comment to column: "type" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."type" IS '抽奖结果';
-- set comment to column: "recharge_amount" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."recharge_amount" IS '充值金额，类型为充值时才有数据';
-- set comment to column: "lotto_id" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."lotto_id" IS '外键：抽奖活动 ID';
-- set comment to column: "user_id" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."user_id" IS '外键：用户 ID';
-- create "prizes" table
CREATE TABLE "prizes" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "level_name" character varying NOT NULL DEFAULT '', "weight" bigint NOT NULL DEFAULT 0, "name" character varying NOT NULL DEFAULT '', "status" character varying NOT NULL DEFAULT 'unknow', "lotto_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "lottoprize_name" to table: "prizes"
CREATE INDEX "lottoprize_name" ON "prizes" ("name");
-- set comment to table: "prizes"
COMMENT ON TABLE "prizes" IS '抽奖活动奖品表';
-- set comment to column: "id" on table: "prizes"
COMMENT ON COLUMN "prizes" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "prizes"
COMMENT ON COLUMN "prizes" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "prizes"
COMMENT ON COLUMN "prizes" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "prizes"
COMMENT ON COLUMN "prizes" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "prizes"
COMMENT ON COLUMN "prizes" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "prizes"
COMMENT ON COLUMN "prizes" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "level_name" on table: "prizes"
COMMENT ON COLUMN "prizes" ."level_name" IS '奖品等级名称';
-- set comment to column: "weight" on table: "prizes"
COMMENT ON COLUMN "prizes" ."weight" IS '奖品等级权重';
-- set comment to column: "name" on table: "prizes"
COMMENT ON COLUMN "prizes" ."name" IS '奖品名称';
-- set comment to column: "status" on table: "prizes"
COMMENT ON COLUMN "prizes" ."status" IS '状态';
-- set comment to column: "lotto_id" on table: "prizes"
COMMENT ON COLUMN "prizes" ."lotto_id" IS '外键：抽奖活动 ID';
-- create "lotto_records" table
CREATE TABLE "lotto_records" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "result" character varying NOT NULL DEFAULT 'unknow', "status" character varying NOT NULL DEFAULT 'waiting', "remain_lotto_count" bigint NOT NULL DEFAULT 0, "lotto_id" bigint NOT NULL DEFAULT 0, "lotto_prize_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "lotto_records"
COMMENT ON TABLE "lotto_records" IS '用户抽奖记录表';
-- set comment to column: "id" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "result" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."result" IS '抽奖结果';
-- set comment to column: "status" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."status" IS '抽奖状态';
-- set comment to column: "remain_lotto_count" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."remain_lotto_count" IS '剩余抽奖次数';
-- set comment to column: "lotto_id" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."lotto_id" IS '外键：抽奖活动 ID';
-- set comment to column: "lotto_prize_id" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."lotto_prize_id" IS '外键：奖品 ID';
-- set comment to column: "user_id" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."user_id" IS '外键：用户 ID';
-- create "lotto_user_counts" table
CREATE TABLE "lotto_user_counts" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "remain_lotto_count" bigint NOT NULL DEFAULT 0, "lotto_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "lotto_user_counts"
COMMENT ON TABLE "lotto_user_counts" IS '用户可用抽奖次数表';
-- set comment to column: "id" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "remain_lotto_count" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."remain_lotto_count" IS '剩余抽奖次数';
-- set comment to column: "lotto_id" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."lotto_id" IS '外键：抽奖活动 ID';
-- set comment to column: "user_id" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."user_id" IS '外键：用户 ID';
