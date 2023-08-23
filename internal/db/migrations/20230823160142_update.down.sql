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
-- reverse: create "vx_socials" table
DROP TABLE "vx_socials";
-- reverse: set comment to column: "user_id" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."user_id" IS '';
-- reverse: set comment to column: "session_key" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."session_key" IS '';
-- reverse: set comment to column: "scope" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."scope" IS '';
-- reverse: set comment to column: "union_id" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."union_id" IS '';
-- reverse: set comment to column: "open_id" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."open_id" IS '';
-- reverse: create "vx_accounts" table
DROP TABLE "vx_accounts";
-- reverse: set comment to column: "user_id" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."user_id" IS '';
-- reverse: set comment to column: "device_id" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."device_id" IS '';
-- reverse: create "user_devices" table
DROP TABLE "user_devices";
-- reverse: set comment to column: "parent_id" on table: "users"
COMMENT ON COLUMN "users" ."parent_id" IS '';
-- reverse: set comment to column: "user_type" on table: "users"
COMMENT ON COLUMN "users" ."user_type" IS '';
-- reverse: set comment to column: "is_frozen" on table: "users"
COMMENT ON COLUMN "users" ."is_frozen" IS '';
-- reverse: set comment to column: "password" on table: "users"
COMMENT ON COLUMN "users" ."password" IS '';
-- reverse: set comment to column: "phone" on table: "users"
COMMENT ON COLUMN "users" ."phone" IS '';
-- reverse: set comment to column: "secret" on table: "users"
COMMENT ON COLUMN "users" ."secret" IS '';
-- reverse: set comment to column: "key" on table: "users"
COMMENT ON COLUMN "users" ."key" IS '';
-- reverse: set comment to column: "jpg_url" on table: "users"
COMMENT ON COLUMN "users" ."jpg_url" IS '';
-- reverse: set comment to column: "name" on table: "users"
COMMENT ON COLUMN "users" ."name" IS '';
-- reverse: create "users" table
DROP TABLE "users";
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
-- reverse: set comment to column: "pure_cep" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."pure_cep" IS '';
-- reverse: set comment to column: "status" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."status" IS '';
-- reverse: create "recharge_orders" table
DROP TABLE "recharge_orders";
-- reverse: set comment to column: "user_id" on table: "profit_settings"
COMMENT ON COLUMN "profit_settings" ."user_id" IS '';
-- reverse: set comment to column: "ratio" on table: "profit_settings"
COMMENT ON COLUMN "profit_settings" ."ratio" IS '';
-- reverse: create "profit_settings" table
DROP TABLE "profit_settings";
-- reverse: set comment to column: "user_id" on table: "profit_accounts"
COMMENT ON COLUMN "profit_accounts" ."user_id" IS '';
-- reverse: set comment to column: "remain_cep" on table: "profit_accounts"
COMMENT ON COLUMN "profit_accounts" ."remain_cep" IS '';
-- reverse: set comment to column: "sum_cep" on table: "profit_accounts"
COMMENT ON COLUMN "profit_accounts" ."sum_cep" IS '';
-- reverse: create index "profit_accounts_user_id_key" to table: "profit_accounts"
DROP INDEX "profit_accounts_user_id_key";
-- reverse: create "profit_accounts" table
DROP TABLE "profit_accounts";
-- reverse: set comment to column: "gift_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."gift_cep" IS '';
-- reverse: set comment to column: "sum_gift_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."sum_gift_cep" IS '';
-- reverse: set comment to column: "pure_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."pure_cep" IS '';
-- reverse: set comment to column: "sum_pure_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."sum_pure_cep" IS '';
-- reverse: set comment to column: "total_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."total_cep" IS '';
-- reverse: set comment to column: "sum_total_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."sum_total_cep" IS '';
-- reverse: create index "platformaccount_type_deleted_at" to table: "platform_accounts"
DROP INDEX "platformaccount_type_deleted_at";
-- reverse: create "platform_accounts" table
DROP TABLE "platform_accounts";
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
-- reverse: create "output_logs" table
DROP TABLE "output_logs";
-- reverse: set comment to column: "user_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."user_id" IS '';
-- reverse: set comment to column: "mission_consume_order_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_consume_order_id" IS '';
-- reverse: set comment to column: "device_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."device_id" IS '';
-- reverse: set comment to column: "serial_number" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."serial_number" IS '';
-- reverse: set comment to column: "is_time" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."is_time" IS '';
-- reverse: set comment to column: "type" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."type" IS '';
-- reverse: set comment to column: "gift_cep" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."gift_cep" IS '';
-- reverse: set comment to column: "pure_cep" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."pure_cep" IS '';
-- reverse: set comment to column: "status" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."status" IS '';
-- reverse: set comment to column: "mission_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_id" IS '';
-- reverse: create "mission_produce_orders" table
DROP TABLE "mission_produce_orders";
-- reverse: set comment to column: "billing_type" on table: "mission_kinds"
COMMENT ON COLUMN "mission_kinds" ."billing_type" IS '';
-- reverse: set comment to column: "category" on table: "mission_kinds"
COMMENT ON COLUMN "mission_kinds" ."category" IS '';
-- reverse: set comment to column: "type" on table: "mission_kinds"
COMMENT ON COLUMN "mission_kinds" ."type" IS '';
-- reverse: create "mission_kinds" table
DROP TABLE "mission_kinds";
-- reverse: set comment to column: "mission_id" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."mission_id" IS '';
-- reverse: set comment to column: "key_pair_id" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."key_pair_id" IS '';
-- reverse: set comment to column: "result_urls" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."result_urls" IS '';
-- reverse: set comment to column: "device_id" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."device_id" IS '';
-- reverse: set comment to column: "result" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."result" IS '';
-- reverse: set comment to column: "finished_at" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."finished_at" IS '';
-- reverse: set comment to column: "started_at" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."started_at" IS '';
-- reverse: create "mission_key_pairs" table
DROP TABLE "mission_key_pairs";
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
-- reverse: set comment to column: "gift_cep" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."gift_cep" IS '';
-- reverse: set comment to column: "pure_cep" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."pure_cep" IS '';
-- reverse: set comment to column: "status" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."status" IS '';
-- reverse: set comment to column: "mission_id" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."mission_id" IS '';
-- reverse: create "mission_consume_orders" table
DROP TABLE "mission_consume_orders";
-- reverse: set comment to column: "user_id" on table: "mission_batches"
COMMENT ON COLUMN "mission_batches" ."user_id" IS '';
-- reverse: set comment to column: "number" on table: "mission_batches"
COMMENT ON COLUMN "mission_batches" ."number" IS '';
-- reverse: create "mission_batches" table
DROP TABLE "mission_batches";
-- reverse: set comment to column: "key_pair_id" on table: "missions"
COMMENT ON COLUMN "missions" ."key_pair_id" IS '';
-- reverse: set comment to column: "mission_batch_number" on table: "missions"
COMMENT ON COLUMN "missions" ."mission_batch_number" IS '';
-- reverse: set comment to column: "result_urls" on table: "missions"
COMMENT ON COLUMN "missions" ."result_urls" IS '';
-- reverse: set comment to column: "result" on table: "missions"
COMMENT ON COLUMN "missions" ."result" IS '';
-- reverse: set comment to column: "status" on table: "missions"
COMMENT ON COLUMN "missions" ."status" IS '';
-- reverse: set comment to column: "call_back_url" on table: "missions"
COMMENT ON COLUMN "missions" ."call_back_url" IS '';
-- reverse: set comment to column: "body" on table: "missions"
COMMENT ON COLUMN "missions" ."body" IS '';
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
-- reverse: create "input_logs" table
DROP TABLE "input_logs";
-- reverse: set comment to column: "caller" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."caller" IS '';
-- reverse: set comment to column: "secret" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."secret" IS '';
-- reverse: set comment to column: "key" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."key" IS '';
-- reverse: create index "hmackeypair_key" to table: "hmac_key_pairs"
DROP INDEX "hmackeypair_key";
-- reverse: create "hmac_key_pairs" table
DROP TABLE "hmac_key_pairs";
-- reverse: set comment to column: "version" on table: "gpus"
COMMENT ON COLUMN "gpus" ."version" IS '';
-- reverse: create "gpus" table
DROP TABLE "gpus";
-- reverse: set comment to column: "mission_status" on table: "enum_mission_status"
COMMENT ON COLUMN "enum_mission_status" ."mission_status" IS '';
-- reverse: set comment to column: "mission_type" on table: "enum_mission_status"
COMMENT ON COLUMN "enum_mission_status" ."mission_type" IS '';
-- reverse: set comment to column: "front_status" on table: "enum_mission_status"
COMMENT ON COLUMN "enum_mission_status" ."front_status" IS '';
-- reverse: create "enum_mission_status" table
DROP TABLE "enum_mission_status";
-- reverse: set comment to column: "mission_call_way" on table: "enum_conditions"
COMMENT ON COLUMN "enum_conditions" ."mission_call_way" IS '';
-- reverse: set comment to column: "mission_type" on table: "enum_conditions"
COMMENT ON COLUMN "enum_conditions" ."mission_type" IS '';
-- reverse: set comment to column: "front_type" on table: "enum_conditions"
COMMENT ON COLUMN "enum_conditions" ."front_type" IS '';
-- reverse: create "enum_conditions" table
DROP TABLE "enum_conditions";
-- reverse: set comment to column: "user_id" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."user_id" IS '';
-- reverse: set comment to column: "profit_account_id" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."profit_account_id" IS '';
-- reverse: set comment to column: "platform_account_id" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."platform_account_id" IS '';
-- reverse: set comment to column: "reason_id" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."reason_id" IS '';
-- reverse: set comment to column: "platform_gift_cep" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."platform_gift_cep" IS '';
-- reverse: set comment to column: "platform_pure_cep" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."platform_pure_cep" IS '';
-- reverse: set comment to column: "gift_cep" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."gift_cep" IS '';
-- reverse: set comment to column: "pure_cep" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."pure_cep" IS '';
-- reverse: set comment to column: "serial_number" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."serial_number" IS '';
-- reverse: set comment to column: "is_minus" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."is_minus" IS '';
-- reverse: set comment to column: "type" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."type" IS '';
-- reverse: create "earn_bills" table
DROP TABLE "earn_bills";
-- reverse: set comment to column: "mission_kind_id" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."mission_kind_id" IS '';
-- reverse: set comment to column: "gpu_id" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."gpu_id" IS '';
-- reverse: set comment to column: "device_id" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."device_id" IS '';
-- reverse: create "device_gpu_missions" table
DROP TABLE "device_gpu_missions";
-- reverse: set comment to column: "user_id" on table: "devices"
COMMENT ON COLUMN "devices" ."user_id" IS '';
-- reverse: set comment to column: "status" on table: "devices"
COMMENT ON COLUMN "devices" ."status" IS '';
-- reverse: set comment to column: "binding_status" on table: "devices"
COMMENT ON COLUMN "devices" ."binding_status" IS '';
-- reverse: set comment to column: "linking" on table: "devices"
COMMENT ON COLUMN "devices" ."linking" IS '';
-- reverse: set comment to column: "sum_cep" on table: "devices"
COMMENT ON COLUMN "devices" ."sum_cep" IS '';
-- reverse: set comment to column: "state" on table: "devices"
COMMENT ON COLUMN "devices" ."state" IS '';
-- reverse: set comment to column: "serial_number" on table: "devices"
COMMENT ON COLUMN "devices" ."serial_number" IS '';
-- reverse: create "devices" table
DROP TABLE "devices";
-- reverse: set comment to column: "user_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."user_id" IS '';
-- reverse: set comment to column: "reason_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."reason_id" IS '';
-- reverse: set comment to column: "cost_account_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."cost_account_id" IS '';
-- reverse: set comment to column: "market_bill_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."market_bill_id" IS '';
-- reverse: set comment to column: "status" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."status" IS '';
-- reverse: set comment to column: "gift_cep" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."gift_cep" IS '';
-- reverse: set comment to column: "pure_cep" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."pure_cep" IS '';
-- reverse: set comment to column: "serial_number" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."serial_number" IS '';
-- reverse: set comment to column: "is_add" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."is_add" IS '';
-- reverse: set comment to column: "type" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."type" IS '';
-- reverse: set comment to table: "cost_bills"
COMMENT ON TABLE "cost_bills" IS '';
-- reverse: create "cost_bills" table
DROP TABLE "cost_bills";
-- reverse: set comment to column: "user_id" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."user_id" IS '';
-- reverse: set comment to column: "frozen_gift_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."frozen_gift_cep" IS '';
-- reverse: set comment to column: "sum_gift_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."sum_gift_cep" IS '';
-- reverse: set comment to column: "gift_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."gift_cep" IS '';
-- reverse: set comment to column: "frozen_pure_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."frozen_pure_cep" IS '';
-- reverse: set comment to column: "sum_pure_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."sum_pure_cep" IS '';
-- reverse: set comment to column: "pure_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."pure_cep" IS '';
-- reverse: set comment to column: "frozen_total_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."frozen_total_cep" IS '';
-- reverse: set comment to column: "sum_total_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."sum_total_cep" IS '';
-- reverse: set comment to column: "total_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."total_cep" IS '';
-- reverse: create index "cost_accounts_user_id_key" to table: "cost_accounts"
DROP INDEX "cost_accounts_user_id_key";
-- reverse: create "cost_accounts" table
DROP TABLE "cost_accounts";
-- reverse: set comment to column: "user_id" on table: "collects"
COMMENT ON COLUMN "collects" ."user_id" IS '';
-- reverse: set comment to column: "jpg_name" on table: "collects"
COMMENT ON COLUMN "collects" ."jpg_name" IS '';
-- reverse: set comment to column: "url" on table: "collects"
COMMENT ON COLUMN "collects" ."url" IS '';
-- reverse: create "collects" table
DROP TABLE "collects";
