-- modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" ALTER COLUMN "business_name" SET DEFAULT '', ALTER COLUMN "business_type" SET DEFAULT 'yun';
-- set comment to column: "business_type" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_type" IS '商户类型';
