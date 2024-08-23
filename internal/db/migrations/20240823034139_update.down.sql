-- reverse: set comment to column: "high_temperature_at" on table: "devices"
COMMENT ON COLUMN "devices" ."high_temperature_at" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "high_temperature_at";
