-- reverse: set comment to column: "type" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."type" IS '抽奖结果';
-- reverse: modify "lotto_get_count_records" table
ALTER TABLE "lotto_get_count_records" ALTER COLUMN "type" SET DEFAULT 'unknow';
