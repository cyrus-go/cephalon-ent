-- reverse: set comment to column: "way" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."way" IS '';
-- reverse: modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" DROP COLUMN "way";
