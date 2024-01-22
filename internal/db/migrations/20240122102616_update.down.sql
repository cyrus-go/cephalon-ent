-- reverse: set comment to column: "highest_earn_month" on table: "gpus"
COMMENT ON COLUMN "gpus" ."highest_earn_month" IS '';
-- reverse: modify "gpus" table
ALTER TABLE "gpus" DROP COLUMN "highest_earn_month";
