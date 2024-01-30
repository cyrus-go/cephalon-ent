-- modify "gpus" table
ALTER TABLE "gpus" ADD COLUMN "highest_earn_month" bigint NOT NULL DEFAULT 0;
-- set comment to column: "highest_earn_month" on table: "gpus"
COMMENT ON COLUMN "gpus" ."highest_earn_month" IS '保底最高月收益';
