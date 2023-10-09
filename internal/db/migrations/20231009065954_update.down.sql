-- reverse: set comment to column: "buy_duration" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."buy_duration" IS '';
-- reverse: modify "mission_orders" table
ALTER TABLE "mission_orders" DROP COLUMN "buy_duration";
