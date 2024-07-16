-- reverse: set comment to column: "stability" on table: "devices"
COMMENT ON COLUMN "devices" ."stability" IS '稳定性，数值越小越稳定';
-- reverse: modify "devices" table
ALTER TABLE "devices" ALTER COLUMN "stability" TYPE bigint, ALTER COLUMN "stability" SET DEFAULT 0;
