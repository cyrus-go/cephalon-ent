-- reverse: set comment to column: "user_id" on table: "login_records"
COMMENT ON COLUMN "login_records" ."user_id" IS '';
-- reverse: set comment to column: "way" on table: "login_records"
COMMENT ON COLUMN "login_records" ."way" IS '';
-- reverse: set comment to column: "ip" on table: "login_records"
COMMENT ON COLUMN "login_records" ."ip" IS '';
-- reverse: set comment to column: "ua" on table: "login_records"
COMMENT ON COLUMN "login_records" ."ua" IS '';
-- reverse: set comment to table: "login_records"
COMMENT ON TABLE "login_records" IS '';
-- reverse: create "login_records" table
DROP TABLE "login_records";
-- reverse: set comment to column: "plan_finished_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."plan_finished_at" IS '';
-- reverse: set comment to column: "plan_started_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."plan_started_at" IS '';
-- reverse: modify "mission_orders" table
ALTER TABLE "mission_orders" DROP COLUMN "plan_finished_at", DROP COLUMN "plan_started_at";
