-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "device_id" character varying NOT NULL DEFAULT '';
-- set comment to column: "device_id" on table: "missions"
COMMENT ON COLUMN "missions" ."device_id" IS '代表该任务是指定设备才能接的任务';
