-- reverse: set comment to column: "last_settled_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."last_settled_at" IS '';
-- reverse: set comment to column: "total_settle_count" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."total_settle_count" IS '';
-- reverse: set comment to column: "settled_count" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."settled_count" IS '';
-- reverse: set comment to column: "settled_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."settled_amount" IS '';
-- reverse: set comment to column: "total_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."total_amount" IS '';
-- reverse: modify "mission_orders" table
ALTER TABLE "mission_orders" DROP COLUMN "last_settled_at", DROP COLUMN "total_settle_count", DROP COLUMN "settled_count", DROP COLUMN "settled_amount", DROP COLUMN "total_amount";
