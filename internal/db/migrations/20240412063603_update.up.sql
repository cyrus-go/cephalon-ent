-- modify "wallets" table
ALTER TABLE "wallets" ADD COLUMN "withdraw_amount" bigint NOT NULL DEFAULT 0;
-- set comment to column: "withdraw_amount" on table: "wallets"
COMMENT ON COLUMN "wallets" ."withdraw_amount" IS '已提现金额，目前只有一种货币可以提现';
