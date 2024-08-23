-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "high_temperature_at" timestamptz NULL;
-- set comment to column: "high_temperature_at" on table: "devices"
COMMENT ON COLUMN "devices" ."high_temperature_at" IS '温度超标的时刻，带时区';
