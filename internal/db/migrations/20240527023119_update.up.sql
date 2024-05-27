-- modify "user_devices" table
ALTER TABLE "user_devices" ADD COLUMN "bind_at" timestamptz NULL, ADD COLUMN "unbind_at" timestamptz NULL;
-- set comment to column: "bind_at" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."bind_at" IS '用户绑定该设备的时间';
-- set comment to column: "unbind_at" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."unbind_at" IS '用户解绑该设备的时间';
