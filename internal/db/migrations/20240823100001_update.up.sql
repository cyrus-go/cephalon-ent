-- modify "device_configs" table
ALTER TABLE "device_configs" ADD COLUMN "gap_random_min" bigint NOT NULL DEFAULT 0;
-- set comment to column: "gap_random_max" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gap_random_max" IS '间隔随机范围最大值';
-- set comment to column: "gap_random_min" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gap_random_min" IS '间隔随机范围最小值';
