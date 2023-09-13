-- create "recharge_campaign_rules" table
CREATE TABLE "recharge_campaign_rules" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "little_value" bigint NOT NULL DEFAULT 0, "large_value" bigint NOT NULL DEFAULT 0, "gift_percent" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "little_value" on table: "recharge_campaign_rules"
COMMENT ON COLUMN "recharge_campaign_rules" ."little_value" IS '充值范围下限';
-- set comment to column: "large_value" on table: "recharge_campaign_rules"
COMMENT ON COLUMN "recharge_campaign_rules" ."large_value" IS '充值范围上限';
-- set comment to column: "gift_percent" on table: "recharge_campaign_rules"
COMMENT ON COLUMN "recharge_campaign_rules" ."gift_percent" IS '赠送的比例';
