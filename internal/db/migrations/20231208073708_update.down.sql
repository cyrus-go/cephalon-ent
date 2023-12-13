-- reverse: set comment to column: "business_type" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_type" IS '商户名称';
-- reverse: modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" ALTER COLUMN "business_type" SET DEFAULT '0', ALTER COLUMN "business_name" SET DEFAULT '0';
