-- reverse: set comment to column: "lowest_earn_month" on table: "gpus"
COMMENT ON COLUMN "gpus" ."lowest_earn_month" IS '';
-- reverse: modify "gpus" table
ALTER TABLE "gpus" DROP COLUMN "lowest_earn_month";
