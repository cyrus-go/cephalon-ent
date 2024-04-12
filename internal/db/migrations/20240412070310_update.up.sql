-- modify "wallets" table
ALTER TABLE "wallets" ADD COLUMN "total_amount" bigint NOT NULL DEFAULT 0;
-- set comment to column: "total_amount" on table: "wallets"
COMMENT ON COLUMN "wallets" ."total_amount" IS '货币总量，当货币是收益货币时，代表总收益，当货币是充值货币时，代表总充值金额';
