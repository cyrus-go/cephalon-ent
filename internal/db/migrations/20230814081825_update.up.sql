-- create index "hmac_key_pairs_user_id_key" to table: "hmac_key_pairs"
CREATE UNIQUE INDEX "hmac_key_pairs_user_id_key" ON "hmac_key_pairs" ("user_id");
-- create index "mission_consume_orders_mission_id_key" to table: "mission_consume_orders"
CREATE UNIQUE INDEX "mission_consume_orders_mission_id_key" ON "mission_consume_orders" ("mission_id");
