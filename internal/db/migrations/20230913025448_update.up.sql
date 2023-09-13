-- modify "cost_bills" table
ALTER TABLE "cost_bills" ALTER COLUMN "type" SET DEFAULT 'unknow', ADD COLUMN "way" character varying NOT NULL DEFAULT 'unknow', ADD COLUMN "market_account_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "way" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."way" IS '额度账户流水的产生方式，微信、支付宝、计时消耗等';
-- set comment to column: "market_account_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."market_account_id" IS '营销账户 id，表示这是一条营销流水（此方案为临时方案）';
