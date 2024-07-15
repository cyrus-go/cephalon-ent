-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "delay" double precision NOT NULL DEFAULT 0, ADD COLUMN "temperature" double precision NOT NULL DEFAULT 0;
-- set comment to column: "delay" on table: "devices"
COMMENT ON COLUMN "devices" ."delay" IS '延迟(单位:ms)';
-- set comment to column: "temperature" on table: "devices"
COMMENT ON COLUMN "devices" ."temperature" IS '温度(单位:℃)';
