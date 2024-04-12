-- reverse: set comment to column: "withdraw_amount" on table: "wallets"
COMMENT ON COLUMN "wallets" ."withdraw_amount" IS '';
-- reverse: modify "wallets" table
ALTER TABLE "wallets" DROP COLUMN "withdraw_amount";
