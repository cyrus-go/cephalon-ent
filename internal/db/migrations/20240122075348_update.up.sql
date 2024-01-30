-- modify "gpus" table
ALTER TABLE "gpus" ADD COLUMN "lowest_earn_month" bigint NOT NULL DEFAULT 0;
-- set comment to column: "lowest_earn_month" on table: "gpus"
COMMENT ON COLUMN "gpus" ."lowest_earn_month" IS '保底最低月收益';
