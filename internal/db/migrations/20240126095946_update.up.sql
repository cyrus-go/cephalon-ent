-- modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" ADD COLUMN "alipay_card_no" character varying NOT NULL DEFAULT '';
-- set comment to column: "alipay_card_no" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."alipay_card_no" IS '支付宝账户';
