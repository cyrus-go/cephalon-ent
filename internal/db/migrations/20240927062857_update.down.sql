-- reverse: set comment to column: "dollar_price" on table: "recharge_campaign_rule_overseas"
COMMENT ON COLUMN "recharge_campaign_rule_overseas" ."dollar_price" IS '美元价格';
-- reverse: modify "recharge_campaign_rule_overseas" table
ALTER TABLE "recharge_campaign_rule_overseas" ALTER COLUMN "original_rmb_price" TYPE double precision, ALTER COLUMN "rmb_price" TYPE double precision, ALTER COLUMN "dollar_price" TYPE double precision;
