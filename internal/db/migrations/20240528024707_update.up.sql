-- modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" ADD COLUMN "current_balance" bigint NOT NULL DEFAULT 0;
-- set comment to column: "current_balance" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."current_balance" IS '当前余额（在生成这条记录时刻的余额），单位：厘';
