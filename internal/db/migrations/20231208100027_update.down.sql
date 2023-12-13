-- reverse: set comment to column: "bank" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."bank" IS '';
-- reverse: set comment to column: "bank_card_number" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."bank_card_number" IS '';
-- reverse: set comment to column: "phone" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."phone" IS '';
-- reverse: set comment to column: "personal_name" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."personal_name" IS '';
-- reverse: set comment to column: "business_id" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_id" IS '货币余额';
-- reverse: set comment to column: "business_name" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_name" IS '商户名称';
-- reverse: modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" DROP COLUMN "bank", DROP COLUMN "bank_card_number", DROP COLUMN "phone", DROP COLUMN "personal_name";
