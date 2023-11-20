-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "cpus" bytea NULL;
-- set comment to column: "cpus" on table: "devices"
COMMENT ON COLUMN "devices" ."cpus" IS 'CPU型号';
