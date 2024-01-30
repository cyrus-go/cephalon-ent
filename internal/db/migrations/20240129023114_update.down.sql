-- reverse: modify "withdraw_accounts" table
ALTER TABLE "withdraw_accounts" ALTER COLUMN "way" SET DEFAULT 'withdraw_bank_card';
