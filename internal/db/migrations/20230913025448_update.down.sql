-- reverse: set comment to column: "market_account_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."market_account_id" IS '';
-- reverse: set comment to column: "way" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."way" IS '';
-- reverse: modify "cost_bills" table
ALTER TABLE "cost_bills" DROP COLUMN "market_account_id", DROP COLUMN "way", ALTER COLUMN "type" SET DEFAULT 'mission';
