-- reverse: set comment to column: "total_amount" on table: "wallets"
COMMENT ON COLUMN "wallets" ."total_amount" IS '';
-- reverse: modify "wallets" table
ALTER TABLE "wallets" DROP COLUMN "total_amount";
