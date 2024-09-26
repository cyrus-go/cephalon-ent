-- create "recharge_campaign_rule_overseas" table
CREATE TABLE "recharge_campaign_rule_overseas" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "dollar_price" double precision NOT NULL DEFAULT 0, "rmb_price" double precision NOT NULL DEFAULT 0, "original_rmb_price" double precision NOT NULL DEFAULT 0, "total_cep" bigint NOT NULL DEFAULT 0, "before_discount_cep" bigint NOT NULL DEFAULT 0, "discount_ratio" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "recharge_campaign_rule_overseas"
COMMENT ON TABLE "recharge_campaign_rule_overseas" IS '海外充值活动的规则，死表，与国内区分，国内是赠送充值比例，海外是固定';
-- set comment to column: "id" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "dollar_price" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."dollar_price" IS '美元价格';
-- set comment to column: "rmb_price" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."rmb_price" IS 'RMB 价格';
-- set comment to column: "original_rmb_price" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."original_rmb_price" IS 'RMB 原价';
-- set comment to column: "total_cep" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."total_cep" IS '总共能获得的脑力值';
-- set comment to column: "before_discount_cep" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."before_discount_cep" IS '优惠前能获得的脑力值';
-- set comment to column: "discount_ratio" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."discount_ratio" IS '优惠力度（现价基于原价的比例）';
