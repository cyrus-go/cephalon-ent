-- reverse: set comment to column: "expired_warning_time" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."expired_warning_time" IS '';
-- reverse: modify "mission_orders" table
ALTER TABLE "mission_orders" DROP COLUMN "expired_warning_time";
