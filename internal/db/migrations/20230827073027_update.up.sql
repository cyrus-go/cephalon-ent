-- modify "gpus" table
ALTER TABLE "gpus" ALTER COLUMN "version" SET DEFAULT 'RTX2060';
-- create index "mission_consume_orders_mission_id_key" to table: "mission_consume_orders"
CREATE UNIQUE INDEX "mission_consume_orders_mission_id_key" ON "mission_consume_orders" ("mission_id");
-- modify "missions" table
ALTER TABLE "missions" ALTER COLUMN "gpu_version" SET DEFAULT 'RTX2060';
-- modify "prices" table
ALTER TABLE "prices" ALTER COLUMN "gpu_version" SET DEFAULT 'RTX2060';
