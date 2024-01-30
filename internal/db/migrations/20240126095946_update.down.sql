-- reverse: set comment to column: "alipay_card_no" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."alipay_card_no" IS '';
-- reverse: modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" DROP COLUMN "alipay_card_no";
