-- modify "mission_load_balances" table
ALTER TABLE "mission_load_balances" ADD COLUMN "state" character varying NOT NULL DEFAULT 'unknown';
-- set comment to column: "state" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."state" IS 'mission 负载均衡总状态';
