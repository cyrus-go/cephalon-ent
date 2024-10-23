-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "last_abnormal_at" timestamptz NULL;
-- set comment to column: "last_abnormal_at" on table: "devices"
COMMENT ON COLUMN "devices" ."last_abnormal_at" IS '最后一次异常时刻，带时区（用于设备稳定性升级）';
