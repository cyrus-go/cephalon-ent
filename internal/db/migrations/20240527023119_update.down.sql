-- reverse: set comment to column: "unbind_at" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."unbind_at" IS '';
-- reverse: set comment to column: "bind_at" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."bind_at" IS '';
-- reverse: modify "user_devices" table
ALTER TABLE "user_devices" DROP COLUMN "unbind_at", DROP COLUMN "bind_at";
