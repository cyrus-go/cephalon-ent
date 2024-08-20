-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "rank_at" timestamptz NULL;
-- set comment to column: "rank_at" on table: "devices"
COMMENT ON COLUMN "devices" ."rank_at" IS '评定设备等级的时刻，带时区';
