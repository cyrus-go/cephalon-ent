-- modify "recharge_campaign_rule_overseas" table
ALTER TABLE "recharge_campaign_rule_overseas" ALTER COLUMN "dollar_price" TYPE character varying, ALTER COLUMN "rmb_price" TYPE character varying, ALTER COLUMN "original_rmb_price" TYPE character varying;
-- set comment to column: "dollar_price" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."dollar_price" IS '美元价格，为了大数计算，都用字符串存';
