-- reverse: modify "prices" table
ALTER TABLE "prices" ALTER COLUMN "gpu_version" SET DEFAULT 'RTX 2060';
-- reverse: modify "missions" table
ALTER TABLE "missions" ALTER COLUMN "gpu_version" SET DEFAULT 'RTX 2060';
-- reverse: create index "mission_consume_orders_mission_id_key" to table: "mission_consume_orders"
DROP INDEX "mission_consume_orders_mission_id_key";
-- reverse: modify "gpus" table
ALTER TABLE "gpus" ALTER COLUMN "version" SET DEFAULT 'RTX 2060';
