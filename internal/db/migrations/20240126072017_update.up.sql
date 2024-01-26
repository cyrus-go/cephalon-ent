-- modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" ADD COLUMN "way" character varying NOT NULL DEFAULT 'withdraw_bank_card';
-- set comment to column: "way" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."way" IS '提现方式';
