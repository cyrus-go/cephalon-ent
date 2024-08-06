-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "fault" character varying NOT NULL DEFAULT 'ok';
-- set comment to column: "fault" on table: "devices"
COMMENT ON COLUMN "devices" ."fault" IS '故障信息';
