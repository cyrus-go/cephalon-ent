-- reverse: set comment to column: "gap_random_min" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gap_random_min" IS '';
-- reverse: set comment to column: "gap_random_max" on table: "device_configs"
COMMENT ON COLUMN "device_configs" ."gap_random_max" IS '间隔随机范围';
-- reverse: modify "device_configs" table
ALTER TABLE "device_configs" DROP COLUMN "gap_random_min";
