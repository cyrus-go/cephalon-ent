-- modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" ADD COLUMN "personal_name" character varying NOT NULL DEFAULT '未设置昵称', ADD COLUMN "phone" character varying NOT NULL DEFAULT '', ADD COLUMN "bank_card_number" character varying NOT NULL DEFAULT '', ADD COLUMN "bank" character varying NOT NULL DEFAULT '未知银行';
-- set comment to column: "business_name" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_name" IS '威付通商户名称';
-- set comment to column: "business_id" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_id" IS '商户 id';
-- set comment to column: "personal_name" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."personal_name" IS '个人商户名称';
-- set comment to column: "phone" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."phone" IS '个人商户手机号';
-- set comment to column: "bank_card_number" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."bank_card_number" IS '银行卡号';
-- set comment to column: "bank" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."bank" IS '开户支行';
