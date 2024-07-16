-- modify "devices" table
ALTER TABLE "devices" ALTER COLUMN "stability" TYPE character varying, ALTER COLUMN "stability" SET DEFAULT 'good';
-- set comment to column: "stability" on table: "devices"
COMMENT ON COLUMN "devices" ."stability" IS '设备稳定性';
