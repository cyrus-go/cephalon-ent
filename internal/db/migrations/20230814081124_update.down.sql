-- reverse: set comment to column: "user_id" on table: "wallets"
COMMENT ON COLUMN "wallets" ."user_id" IS '';
-- reverse: set comment to column: "sum_cep" on table: "wallets"
COMMENT ON COLUMN "wallets" ."sum_cep" IS '';
-- reverse: set comment to column: "cep" on table: "wallets"
COMMENT ON COLUMN "wallets" ."cep" IS '';
-- reverse: set comment to table: "wallets"
COMMENT ON TABLE "wallets" IS '';
-- reverse: create index "wallets_user_id_key" to table: "wallets"
DROP INDEX "wallets_user_id_key";
-- reverse: create "wallets" table
DROP TABLE "wallets";
-- reverse: set comment to column: "user_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."user_id" IS '';
-- reverse: set comment to column: "refresh_token" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."refresh_token" IS '';
-- reverse: set comment to column: "access_token" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."access_token" IS '';
-- reverse: set comment to column: "session_key" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."session_key" IS '';
-- reverse: set comment to column: "scope" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."scope" IS '';
-- reverse: set comment to column: "union_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."union_id" IS '';
-- reverse: set comment to column: "open_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."open_id" IS '';
-- reverse: set comment to column: "app_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."app_id" IS '';
-- reverse: set comment to table: "vx_socials"
COMMENT ON TABLE "vx_socials" IS '';
-- reverse: create "vx_socials" table
DROP TABLE "vx_socials";
-- reverse: set comment to column: "user_id" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."user_id" IS '';
-- reverse: set comment to column: "device_id" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."device_id" IS '';
-- reverse: set comment to table: "user_devices"
COMMENT ON TABLE "user_devices" IS '';
-- reverse: create "user_devices" table
DROP TABLE "user_devices";
-- reverse: set comment to column: "social_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."social_id" IS '';
-- reverse: set comment to column: "user_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."user_id" IS '';
-- reverse: set comment to column: "out_transaction_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."out_transaction_id" IS '';
-- reverse: set comment to column: "from_user_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."from_user_id" IS '';
-- reverse: set comment to column: "third_api_resp" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."third_api_resp" IS '';
-- reverse: set comment to column: "serial_number" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."serial_number" IS '';
-- reverse: set comment to column: "type" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."type" IS '';
-- reverse: set comment to column: "cep" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."cep" IS '';
-- reverse: set comment to column: "status" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."status" IS '';
-- reverse: set comment to table: "recharge_orders"
COMMENT ON TABLE "recharge_orders" IS '';
-- reverse: create "recharge_orders" table
DROP TABLE "recharge_orders";
-- reverse: set comment to column: "user_id" on table: "profit_settings"
COMMENT ON COLUMN "profit_settings" ."user_id" IS '';
-- reverse: set comment to column: "ratio" on table: "profit_settings"
COMMENT ON COLUMN "profit_settings" ."ratio" IS '';
-- reverse: set comment to table: "profit_settings"
COMMENT ON TABLE "profit_settings" IS '';
-- reverse: create "profit_settings" table
DROP TABLE "profit_settings";
-- reverse: set comment to column: "cep" on table: "platform_wallets"
COMMENT ON COLUMN "platform_wallets" ."cep" IS '';
-- reverse: set comment to column: "sum_cep" on table: "platform_wallets"
COMMENT ON COLUMN "platform_wallets" ."sum_cep" IS '';
-- reverse: set comment to table: "platform_wallets"
COMMENT ON TABLE "platform_wallets" IS '';
-- reverse: create index "platformwallet_type_deleted_at" to table: "platform_wallets"
DROP INDEX "platformwallet_type_deleted_at";
-- reverse: create "platform_wallets" table
DROP TABLE "platform_wallets";
-- reverse: set comment to column: "hmac_key" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."hmac_key" IS '';
-- reverse: set comment to column: "status" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."status" IS '';
-- reverse: set comment to column: "caller" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."caller" IS '';
-- reverse: set comment to column: "ip" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."ip" IS '';
-- reverse: set comment to column: "url" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."url" IS '';
-- reverse: set comment to column: "body" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."body" IS '';
-- reverse: set comment to column: "headers" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."headers" IS '';
-- reverse: set comment to column: "trace_id" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."trace_id" IS '';
-- reverse: set comment to table: "output_logs"
COMMENT ON TABLE "output_logs" IS '';
-- reverse: create "output_logs" table
DROP TABLE "output_logs";
-- reverse: set comment to column: "category" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."category" IS '';
-- reverse: set comment to column: "is_time" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."is_time" IS '';
-- reverse: set comment to column: "cep" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."cep" IS '';
-- reverse: set comment to column: "gpu" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."gpu" IS '';
-- reverse: set comment to column: "type" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."type" IS '';
-- reverse: set comment to table: "mission_types"
COMMENT ON TABLE "mission_types" IS '';
-- reverse: create "mission_types" table
DROP TABLE "mission_types";
-- reverse: set comment to column: "mission_id" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."mission_id" IS '';
-- reverse: set comment to column: "hmac_key_pair_id" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."hmac_key_pair_id" IS '';
-- reverse: set comment to column: "device_id" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."device_id" IS '';
-- reverse: set comment to column: "additional_result" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."additional_result" IS '';
-- reverse: set comment to column: "result_urls" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."result_urls" IS '';
-- reverse: set comment to column: "status" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."status" IS '';
-- reverse: set comment to column: "finished_at" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."finished_at" IS '';
-- reverse: set comment to column: "started_at" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."started_at" IS '';
-- reverse: set comment to table: "mission_productions"
COMMENT ON TABLE "mission_productions" IS '';
-- reverse: create "mission_productions" table
DROP TABLE "mission_productions";
-- reverse: set comment to column: "user_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."user_id" IS '';
-- reverse: set comment to column: "mission_production_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_production_id" IS '';
-- reverse: set comment to column: "mission_consume_order_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_consume_order_id" IS '';
-- reverse: set comment to column: "mission_batch_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_batch_id" IS '';
-- reverse: set comment to column: "mission_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_id" IS '';
-- reverse: set comment to column: "device_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."device_id" IS '';
-- reverse: set comment to column: "serial_number" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."serial_number" IS '';
-- reverse: set comment to column: "is_time" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."is_time" IS '';
-- reverse: set comment to column: "type" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."type" IS '';
-- reverse: set comment to column: "cep" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."cep" IS '';
-- reverse: set comment to column: "status" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."status" IS '';
-- reverse: set comment to table: "mission_produce_orders"
COMMENT ON TABLE "mission_produce_orders" IS '';
-- reverse: create index "mission_produce_orders_mission_production_id_key" to table: "mission_produce_orders"
DROP INDEX "mission_produce_orders_mission_production_id_key";
-- reverse: create "mission_produce_orders" table
DROP TABLE "mission_produce_orders";
-- reverse: set comment to column: "user_id" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."user_id" IS '';
-- reverse: set comment to column: "mission_batch_id" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."mission_batch_id" IS '';
-- reverse: set comment to column: "mission_batch_number" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."mission_batch_number" IS '';
-- reverse: set comment to column: "finished_at" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."finished_at" IS '';
-- reverse: set comment to column: "started_at" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."started_at" IS '';
-- reverse: set comment to column: "serial_number" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."serial_number" IS '';
-- reverse: set comment to column: "call_way" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."call_way" IS '';
-- reverse: set comment to column: "is_time" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."is_time" IS '';
-- reverse: set comment to column: "type" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."type" IS '';
-- reverse: set comment to column: "cep" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."cep" IS '';
-- reverse: set comment to column: "status" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."status" IS '';
-- reverse: set comment to column: "mission_id" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."mission_id" IS '';
-- reverse: set comment to table: "mission_consume_orders"
COMMENT ON TABLE "mission_consume_orders" IS '';
-- reverse: create index "mission_consume_orders_mission_mission_consume_order_key" to table: "mission_consume_orders"
DROP INDEX "mission_consume_orders_mission_mission_consume_order_key";
-- reverse: create "mission_consume_orders" table
DROP TABLE "mission_consume_orders";
-- reverse: set comment to column: "user_id" on table: "mission_batches"
COMMENT ON COLUMN "mission_batches" ."user_id" IS '';
-- reverse: set comment to column: "number" on table: "mission_batches"
COMMENT ON COLUMN "mission_batches" ."number" IS '';
-- reverse: set comment to table: "mission_batches"
COMMENT ON TABLE "mission_batches" IS '';
-- reverse: create "mission_batches" table
DROP TABLE "mission_batches";
-- reverse: set comment to column: "user_id" on table: "missions"
COMMENT ON COLUMN "missions" ."user_id" IS '';
-- reverse: set comment to column: "mission_batch_id" on table: "missions"
COMMENT ON COLUMN "missions" ."mission_batch_id" IS '';
-- reverse: set comment to column: "hmac_key_pair_id" on table: "missions"
COMMENT ON COLUMN "missions" ."hmac_key_pair_id" IS '';
-- reverse: set comment to column: "mission_batch_number" on table: "missions"
COMMENT ON COLUMN "missions" ."mission_batch_number" IS '';
-- reverse: set comment to column: "additional_result" on table: "missions"
COMMENT ON COLUMN "missions" ."additional_result" IS '';
-- reverse: set comment to column: "result_urls" on table: "missions"
COMMENT ON COLUMN "missions" ."result_urls" IS '';
-- reverse: set comment to column: "status" on table: "missions"
COMMENT ON COLUMN "missions" ."status" IS '';
-- reverse: set comment to column: "call_back_url" on table: "missions"
COMMENT ON COLUMN "missions" ."call_back_url" IS '';
-- reverse: set comment to column: "body" on table: "missions"
COMMENT ON COLUMN "missions" ."body" IS '';
-- reverse: set comment to column: "is_time" on table: "missions"
COMMENT ON COLUMN "missions" ."is_time" IS '';
-- reverse: set comment to column: "type" on table: "missions"
COMMENT ON COLUMN "missions" ."type" IS '';
-- reverse: create "missions" table
DROP TABLE "missions";
-- reverse: set comment to column: "hmac_key" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."hmac_key" IS '';
-- reverse: set comment to column: "method" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."method" IS '';
-- reverse: set comment to column: "caller" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."caller" IS '';
-- reverse: set comment to column: "ip" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."ip" IS '';
-- reverse: set comment to column: "url" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."url" IS '';
-- reverse: set comment to column: "query" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."query" IS '';
-- reverse: set comment to column: "body" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."body" IS '';
-- reverse: set comment to column: "headers" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."headers" IS '';
-- reverse: set comment to column: "trace_id" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."trace_id" IS '';
-- reverse: set comment to table: "input_logs"
COMMENT ON TABLE "input_logs" IS '';
-- reverse: create "input_logs" table
DROP TABLE "input_logs";
-- reverse: set comment to column: "user_id" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."user_id" IS '';
-- reverse: set comment to column: "caller" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."caller" IS '';
-- reverse: set comment to column: "secret" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."secret" IS '';
-- reverse: set comment to column: "key" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."key" IS '';
-- reverse: set comment to table: "hmac_key_pairs"
COMMENT ON TABLE "hmac_key_pairs" IS '';
-- reverse: create index "hmackeypair_key" to table: "hmac_key_pairs"
DROP INDEX "hmackeypair_key";
-- reverse: create index "hmac_key_pairs_user_hmac_key_pair_key" to table: "hmac_key_pairs"
DROP INDEX "hmac_key_pairs_user_hmac_key_pair_key";
-- reverse: create "hmac_key_pairs" table
DROP TABLE "hmac_key_pairs";
-- reverse: set comment to column: "user_id" on table: "devices"
COMMENT ON COLUMN "devices" ."user_id" IS '';
-- reverse: set comment to column: "binding_status" on table: "devices"
COMMENT ON COLUMN "devices" ."binding_status" IS '';
-- reverse: set comment to column: "status" on table: "devices"
COMMENT ON COLUMN "devices" ."status" IS '';
-- reverse: set comment to table: "devices"
COMMENT ON TABLE "devices" IS '';
-- reverse: create "devices" table
DROP TABLE "devices";
-- reverse: set comment to column: "user_id" on table: "collections"
COMMENT ON COLUMN "collections" ."user_id" IS '';
-- reverse: set comment to column: "picture_name" on table: "collections"
COMMENT ON COLUMN "collections" ."picture_name" IS '';
-- reverse: set comment to column: "url" on table: "collections"
COMMENT ON COLUMN "collections" ."url" IS '';
-- reverse: set comment to table: "collections"
COMMENT ON TABLE "collections" IS '';
-- reverse: create "collections" table
DROP TABLE "collections";
-- reverse: set comment to column: "wallet_id" on table: "bills"
COMMENT ON COLUMN "bills" ."wallet_id" IS '';
-- reverse: set comment to column: "user_id" on table: "bills"
COMMENT ON COLUMN "bills" ."user_id" IS '';
-- reverse: set comment to column: "platform_wallet_id" on table: "bills"
COMMENT ON COLUMN "bills" ."platform_wallet_id" IS '';
-- reverse: set comment to column: "reason_id" on table: "bills"
COMMENT ON COLUMN "bills" ."reason_id" IS '';
-- reverse: set comment to column: "platform_cep" on table: "bills"
COMMENT ON COLUMN "bills" ."platform_cep" IS '';
-- reverse: set comment to column: "market_bill_id" on table: "bills"
COMMENT ON COLUMN "bills" ."market_bill_id" IS '';
-- reverse: set comment to column: "status" on table: "bills"
COMMENT ON COLUMN "bills" ."status" IS '';
-- reverse: set comment to column: "cep" on table: "bills"
COMMENT ON COLUMN "bills" ."cep" IS '';
-- reverse: set comment to column: "serial_number" on table: "bills"
COMMENT ON COLUMN "bills" ."serial_number" IS '';
-- reverse: set comment to column: "is_add" on table: "bills"
COMMENT ON COLUMN "bills" ."is_add" IS '';
-- reverse: set comment to column: "type" on table: "bills"
COMMENT ON COLUMN "bills" ."type" IS '';
-- reverse: set comment to table: "bills"
COMMENT ON TABLE "bills" IS '';
-- reverse: create "bills" table
DROP TABLE "bills";
-- reverse: set comment to table: "users"
COMMENT ON TABLE "users" IS '用户表，手机号唯一';
