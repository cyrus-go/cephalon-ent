-- reverse: set comment to column: "id_card" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."id_card" IS '';
-- reverse: modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" DROP COLUMN "id_card";
