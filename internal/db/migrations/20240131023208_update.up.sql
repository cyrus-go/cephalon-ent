-- create "lotto_chance_rules" table
CREATE TABLE "lotto_chance_rules" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "condition" character varying NOT NULL DEFAULT 'unknown', "award_count" bigint NOT NULL DEFAULT 0, "recharge_amount" bigint NOT NULL DEFAULT 0, "lotto_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "lotto_chance_rules"
COMMENT ON TABLE "lotto_chance_rules" IS '获取抽奖次数规则表';
-- set comment to column: "id" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "condition" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."condition" IS '条件';
-- set comment to column: "award_count" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."award_count" IS '奖励抽奖次数';
-- set comment to column: "recharge_amount" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."recharge_amount" IS '充值金额，类型为充值时才有数据';
-- set comment to column: "lotto_id" on table: "lotto_chance_rules"
COMMENT ON COLUMN "lotto_chance_rules" ."lotto_id" IS '外键：抽奖活动 ID';
