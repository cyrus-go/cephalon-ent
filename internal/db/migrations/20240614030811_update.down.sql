-- reverse: set comment to column: "company_account" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."company_account" IS '';
-- reverse: set comment to column: "business_name" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_name" IS '威付通商户名称';
-- reverse: modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" DROP COLUMN "company_account";
