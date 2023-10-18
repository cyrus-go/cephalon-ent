-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "white_device_ids" bytea NULL, ADD COLUMN "black_device_ids" bytea NULL;
-- set comment to column: "white_device_ids" on table: "missions"
COMMENT ON COLUMN "missions" ."white_device_ids" IS '任务的设备白名单';
-- set comment to column: "black_device_ids" on table: "missions"
COMMENT ON COLUMN "missions" ."black_device_ids" IS '任务的设备黑名单';
