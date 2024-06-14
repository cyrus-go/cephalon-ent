-- modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" ADD COLUMN "company_account" character varying NOT NULL DEFAULT '';
-- set comment to column: "business_name" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_name" IS '威付通商户名称，对公时为户名';
-- set comment to column: "company_account" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."company_account" IS '对公账号';
