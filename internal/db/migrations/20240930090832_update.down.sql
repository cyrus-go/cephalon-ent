-- reverse: set comment to column: "access_count" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."access_count" IS '';
-- reverse: set comment to column: "last_access" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."last_access" IS '';
-- reverse: set comment to column: "mission_load_balance_id" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."mission_load_balance_id" IS '';
-- reverse: set comment to column: "mission_id" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."mission_id" IS '';
-- reverse: set comment to column: "deleted_at" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "mission_load_balance_accesses"
COMMENT ON COLUMN "mission_load_balance_accesses" ."id" IS '';
-- reverse: set comment to table: "mission_load_balance_accesses"
COMMENT ON TABLE "mission_load_balance_accesses" IS '';
-- reverse: create index "missionloadbalanceaccess_mission_load_balance_id" to table: "mission_load_balance_accesses"
DROP INDEX "missionloadbalanceaccess_mission_load_balance_id";
-- reverse: create index "missionloadbalanceaccess_mission_id" to table: "mission_load_balance_accesses"
DROP INDEX "missionloadbalanceaccess_mission_id";
-- reverse: create "mission_load_balance_accesses" table
DROP TABLE "mission_load_balance_accesses";
-- reverse: set comment to column: "mission_batch_number" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."mission_batch_number" IS '';
-- reverse: set comment to column: "mission_batch_id" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."mission_batch_id" IS '';
-- reverse: set comment to column: "min_mission_count" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."min_mission_count" IS '';
-- reverse: set comment to column: "max_mission_count" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."max_mission_count" IS '';
-- reverse: set comment to column: "gpu_num" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."gpu_num" IS '';
-- reverse: set comment to column: "gpu_version" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."gpu_version" IS '';
-- reverse: set comment to column: "finished_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."finished_at" IS '';
-- reverse: set comment to column: "started_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."started_at" IS '';
-- reverse: set comment to column: "user_id" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."user_id" IS '';
-- reverse: set comment to column: "mission_type" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."mission_type" IS '';
-- reverse: set comment to column: "deleted_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."id" IS '';
-- reverse: set comment to table: "mission_load_balances"
COMMENT ON TABLE "mission_load_balances" IS '';
-- reverse: create index "missionloadbalance_mission_batch_number" to table: "mission_load_balances"
DROP INDEX "missionloadbalance_mission_batch_number";
-- reverse: create index "missionloadbalance_mission_batch_id" to table: "mission_load_balances"
DROP INDEX "missionloadbalance_mission_batch_id";
-- reverse: create "mission_load_balances" table
DROP TABLE "mission_load_balances";
