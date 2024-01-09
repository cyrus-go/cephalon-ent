-- modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" ADD COLUMN "id_card" character varying NOT NULL DEFAULT '';
-- set comment to column: "id_card" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."id_card" IS '身份证号码';
