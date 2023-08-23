-- create "collects" table
CREATE TABLE "collects" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "url" character varying NOT NULL DEFAULT '', "jpg_name" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "url" on table: "collects"
COMMENT ON COLUMN "collects" ."url" IS '路径';
-- set comment to column: "jpg_name" on table: "collects"
COMMENT ON COLUMN "collects" ."jpg_name" IS '照片名字';
-- set comment to column: "user_id" on table: "collects"
COMMENT ON COLUMN "collects" ."user_id" IS '外键用户 id';
-- create "cost_accounts" table
CREATE TABLE "cost_accounts" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "total_cep" bigint NOT NULL DEFAULT 0, "sum_total_cep" bigint NOT NULL DEFAULT 0, "frozen_total_cep" bigint NOT NULL DEFAULT 0, "pure_cep" bigint NOT NULL DEFAULT 0, "sum_pure_cep" bigint NOT NULL DEFAULT 0, "frozen_pure_cep" bigint NOT NULL DEFAULT 0, "gift_cep" bigint NOT NULL DEFAULT 0, "sum_gift_cep" bigint NOT NULL DEFAULT 0, "frozen_gift_cep" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "cost_accounts_user_id_key" to table: "cost_accounts"
CREATE UNIQUE INDEX "cost_accounts_user_id_key" ON "cost_accounts" ("user_id");
-- set comment to column: "total_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."total_cep" IS '总余额';
-- set comment to column: "sum_total_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."sum_total_cep" IS '累计总余额';
-- set comment to column: "frozen_total_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."frozen_total_cep" IS '暂时冻结的总余额';
-- set comment to column: "pure_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."pure_cep" IS '本金余额';
-- set comment to column: "sum_pure_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."sum_pure_cep" IS '累计本金余额';
-- set comment to column: "frozen_pure_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."frozen_pure_cep" IS '暂时冻结的本金余额';
-- set comment to column: "gift_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."gift_cep" IS '赠送余额';
-- set comment to column: "sum_gift_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."sum_gift_cep" IS '累计赠送余额';
-- set comment to column: "frozen_gift_cep" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."frozen_gift_cep" IS '暂时冻结的赠金余额';
-- set comment to column: "user_id" on table: "cost_accounts"
COMMENT ON COLUMN "cost_accounts" ."user_id" IS '外键用户 id';
-- create "cost_bills" table
CREATE TABLE "cost_bills" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'mission', "is_add" boolean NOT NULL DEFAULT false, "serial_number" character varying NOT NULL DEFAULT '', "pure_cep" bigint NOT NULL DEFAULT 0, "gift_cep" bigint NOT NULL DEFAULT 0, "status" character varying NOT NULL DEFAULT 'pending', "market_bill_id" bigint NOT NULL DEFAULT 0, "cost_account_id" bigint NOT NULL DEFAULT 0, "reason_id" bigint NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "cost_bills"
COMMENT ON TABLE "cost_bills" IS '记录额度账户的变动的主流水账单记录';
-- set comment to column: "type" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."type" IS '额度账户流水的类型，充值或者任务消耗';
-- set comment to column: "is_add" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."is_add" IS '是否增加余额，布尔值默认为否';
-- set comment to column: "serial_number" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."serial_number" IS '账单序列号';
-- set comment to column: "pure_cep" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."pure_cep" IS '消耗多少本金余额';
-- set comment to column: "gift_cep" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."gift_cep" IS '消耗多少赠送余额';
-- set comment to column: "status" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."status" IS '消耗流水一开始不能直接生效，确定关联的消耗时间成功后才可以扣费';
-- set comment to column: "market_bill_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."market_bill_id" IS '营销流水 id';
-- set comment to column: "cost_account_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."cost_account_id" IS '外键消耗账户 id';
-- set comment to column: "reason_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."reason_id" IS '关联消耗产生的来源外键 id，比如 type 为 mission 时关联任务订单';
-- set comment to column: "user_id" on table: "cost_bills"
COMMENT ON COLUMN "cost_bills" ."user_id" IS '外键用户 id';
-- create "devices" table
CREATE TABLE "devices" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "serial_number" character varying NOT NULL DEFAULT '', "state" character varying NOT NULL DEFAULT 'Init', "sum_cep" bigint NOT NULL DEFAULT 0, "linking" boolean NOT NULL DEFAULT false, "binding_status" character varying NOT NULL DEFAULT 'init', "status" character varying NOT NULL DEFAULT 'online', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "serial_number" on table: "devices"
COMMENT ON COLUMN "devices" ."serial_number" IS '设备唯一序列号';
-- set comment to column: "state" on table: "devices"
COMMENT ON COLUMN "devices" ."state" IS '设备状态';
-- set comment to column: "sum_cep" on table: "devices"
COMMENT ON COLUMN "devices" ."sum_cep" IS '该设备总获得利润';
-- set comment to column: "linking" on table: "devices"
COMMENT ON COLUMN "devices" ."linking" IS '设备是否正在对接中';
-- set comment to column: "binding_status" on table: "devices"
COMMENT ON COLUMN "devices" ."binding_status" IS '设备的绑定状态';
-- set comment to column: "status" on table: "devices"
COMMENT ON COLUMN "devices" ."status" IS '设备状态';
-- set comment to column: "user_id" on table: "devices"
COMMENT ON COLUMN "devices" ."user_id" IS '外键用户 id';
-- create "device_gpu_missions" table
CREATE TABLE "device_gpu_missions" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "device_id" bigint NOT NULL DEFAULT 0, "gpu_id" bigint NOT NULL DEFAULT 0, "mission_kind_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "device_id" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."device_id" IS '外键设备 id';
-- set comment to column: "gpu_id" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."gpu_id" IS '外键 gpu id';
-- set comment to column: "mission_kind_id" on table: "device_gpu_missions"
COMMENT ON COLUMN "device_gpu_missions" ."mission_kind_id" IS '外键任务种类 id';
-- create "earn_bills" table
CREATE TABLE "earn_bills" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'mission', "is_minus" boolean NOT NULL DEFAULT false, "serial_number" character varying NOT NULL DEFAULT '', "pure_cep" bigint NOT NULL DEFAULT 0, "gift_cep" bigint NOT NULL DEFAULT 0, "platform_pure_cep" bigint NOT NULL DEFAULT 0, "platform_gift_cep" bigint NOT NULL DEFAULT 0, "reason_id" bigint NULL DEFAULT 0, "platform_account_id" bigint NOT NULL DEFAULT 0, "profit_account_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "type" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."type" IS '分润账户变动类型';
-- set comment to column: "is_minus" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."is_minus" IS '是否减少分润钱包余额，默认为否';
-- set comment to column: "serial_number" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."serial_number" IS '账单序列号';
-- set comment to column: "pure_cep" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."pure_cep" IS '分润多少本金余额';
-- set comment to column: "gift_cep" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."gift_cep" IS '分润多少赠送余额';
-- set comment to column: "platform_pure_cep" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."platform_pure_cep" IS '平台收取多少本金余额';
-- set comment to column: "platform_gift_cep" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."platform_gift_cep" IS '平台收取多少赠送余额';
-- set comment to column: "reason_id" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."reason_id" IS '关联分润产生的来源外键 id，比如 type 为 mission 时关联任务订单';
-- set comment to column: "platform_account_id" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."platform_account_id" IS '平台分润账户 id';
-- set comment to column: "profit_account_id" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."profit_account_id" IS '外键分润账户 id';
-- set comment to column: "user_id" on table: "earn_bills"
COMMENT ON COLUMN "earn_bills" ."user_id" IS '外键用户 id';
-- create "enum_conditions" table
CREATE TABLE "enum_conditions" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "front_type" character varying NOT NULL DEFAULT '', "mission_type" character varying NOT NULL DEFAULT '', "mission_call_way" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- set comment to column: "front_type" on table: "enum_conditions"
COMMENT ON COLUMN "enum_conditions" ."front_type" IS '给到前端的类型';
-- set comment to column: "mission_type" on table: "enum_conditions"
COMMENT ON COLUMN "enum_conditions" ."mission_type" IS '任务类型';
-- set comment to column: "mission_call_way" on table: "enum_conditions"
COMMENT ON COLUMN "enum_conditions" ."mission_call_way" IS '任务调用方式';
-- create "enum_mission_status" table
CREATE TABLE "enum_mission_status" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "front_status" character varying NOT NULL DEFAULT '', "mission_type" character varying NOT NULL DEFAULT '', "mission_status" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- set comment to column: "front_status" on table: "enum_mission_status"
COMMENT ON COLUMN "enum_mission_status" ."front_status" IS '给到前端的任务状态';
-- set comment to column: "mission_type" on table: "enum_mission_status"
COMMENT ON COLUMN "enum_mission_status" ."mission_type" IS '任务类型';
-- set comment to column: "mission_status" on table: "enum_mission_status"
COMMENT ON COLUMN "enum_mission_status" ."mission_status" IS '任务状态';
-- create "gpus" table
CREATE TABLE "gpus" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "version" character varying NOT NULL DEFAULT 'RTX 2060', PRIMARY KEY ("id"));
-- set comment to column: "version" on table: "gpus"
COMMENT ON COLUMN "gpus" ."version" IS '显卡型号';
-- create "hmac_key_pairs" table
CREATE TABLE "hmac_key_pairs" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "key" character varying NOT NULL, "secret" character varying NOT NULL, "caller" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- create index "hmackeypair_key" to table: "hmac_key_pairs"
CREATE INDEX "hmackeypair_key" ON "hmac_key_pairs" ("key");
-- set comment to column: "key" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."key" IS '密钥对的 key 值，用于检索密钥';
-- set comment to column: "secret" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."secret" IS '加密密钥';
-- set comment to column: "caller" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."caller" IS '请求方';
-- create "input_logs" table
CREATE TABLE "input_logs" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "trace_id" bigint NOT NULL, "headers" character varying NOT NULL, "body" character varying NULL DEFAULT '', "query" character varying NULL DEFAULT '', "url" character varying NOT NULL, "ip" character varying NULL DEFAULT '', "caller" character varying NOT NULL DEFAULT 'general', "method" character varying NOT NULL, "hmac_key" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- set comment to column: "trace_id" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."trace_id" IS '请求追踪 id';
-- set comment to column: "headers" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."headers" IS '请求头';
-- set comment to column: "body" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."body" IS '请求体';
-- set comment to column: "query" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."query" IS 'Query 参数';
-- set comment to column: "url" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."url" IS '请求地址';
-- set comment to column: "ip" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."ip" IS '客户端 IP';
-- set comment to column: "caller" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."caller" IS '调用方';
-- set comment to column: "method" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."method" IS '请求方式';
-- set comment to column: "hmac_key" on table: "input_logs"
COMMENT ON COLUMN "input_logs" ."hmac_key" IS '记录调用者的密钥对';
-- create "missions" table
CREATE TABLE "missions" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'txt2img', "body" character varying NOT NULL DEFAULT '', "call_back_url" character varying NOT NULL DEFAULT '', "status" character varying NOT NULL DEFAULT 'waiting', "result" character varying NOT NULL DEFAULT 'pending', "result_urls" jsonb NULL, "mission_batch_number" character varying NOT NULL DEFAULT '', "key_pair_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "type" on table: "missions"
COMMENT ON COLUMN "missions" ."type" IS '任务类型';
-- set comment to column: "body" on table: "missions"
COMMENT ON COLUMN "missions" ."body" IS '任务的请求参数体';
-- set comment to column: "call_back_url" on table: "missions"
COMMENT ON COLUMN "missions" ."call_back_url" IS '回调地址，空字符串表示不回调';
-- set comment to column: "status" on table: "missions"
COMMENT ON COLUMN "missions" ."status" IS '任务状态';
-- set comment to column: "result" on table: "missions"
COMMENT ON COLUMN "missions" ."result" IS '任务结果，pending 表示还没有结果';
-- set comment to column: "result_urls" on table: "missions"
COMMENT ON COLUMN "missions" ."result_urls" IS '任务结果资源位置列表序列化';
-- set comment to column: "mission_batch_number" on table: "missions"
COMMENT ON COLUMN "missions" ."mission_batch_number" IS '任务批次号';
-- set comment to column: "key_pair_id" on table: "missions"
COMMENT ON COLUMN "missions" ."key_pair_id" IS '任务创建者的密钥对 ID';
-- create "mission_batches" table
CREATE TABLE "mission_batches" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "number" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "number" on table: "mission_batches"
COMMENT ON COLUMN "mission_batches" ."number" IS '批次号码';
-- set comment to column: "user_id" on table: "mission_batches"
COMMENT ON COLUMN "mission_batches" ."user_id" IS '创建者';
-- create "mission_consume_orders" table
CREATE TABLE "mission_consume_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "mission_id" bigint NOT NULL DEFAULT 0, "status" character varying NOT NULL DEFAULT 'waiting', "pure_cep" bigint NOT NULL DEFAULT 0, "gift_cep" bigint NOT NULL DEFAULT 0, "type" character varying NOT NULL DEFAULT 'txt2img', "is_time" boolean NOT NULL DEFAULT false, "call_way" character varying NOT NULL DEFAULT 'api', "serial_number" character varying NOT NULL DEFAULT '', "started_at" timestamptz NOT NULL, "finished_at" timestamptz NOT NULL, "mission_batch_number" character varying NOT NULL DEFAULT '', "mission_batch_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "mission_id" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."mission_id" IS '任务 id，关联任务中枢的任务';
-- set comment to column: "status" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."status" IS '任务订单的状态，注意不强关联任务的状态';
-- set comment to column: "pure_cep" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."pure_cep" IS '任务消耗的本金 cep 量';
-- set comment to column: "gift_cep" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."gift_cep" IS '任务消耗的赠送 cep 量';
-- set comment to column: "type" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."type" IS '任务类型，等于任务表的类型字段';
-- set comment to column: "is_time" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."is_time" IS '是否为计时类型任务';
-- set comment to column: "call_way" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."call_way" IS '调用方式，API 调用或者微信小程序调用';
-- set comment to column: "serial_number" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."serial_number" IS '订单序列号';
-- set comment to column: "started_at" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."started_at" IS '任务开始执行时刻';
-- set comment to column: "finished_at" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."finished_at" IS '任务结束执行时刻';
-- set comment to column: "mission_batch_number" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."mission_batch_number" IS '任务批次号，用于方便检索';
-- set comment to column: "mission_batch_id" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."mission_batch_id" IS '任务批次外键';
-- set comment to column: "user_id" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."user_id" IS '外键关联用户 id';
-- create "mission_key_pairs" table
CREATE TABLE "mission_key_pairs" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "started_at" timestamptz NOT NULL, "finished_at" timestamptz NOT NULL, "result" character varying NOT NULL DEFAULT 'pending', "device_id" bigint NOT NULL DEFAULT 0, "result_urls" jsonb NULL, "key_pair_id" bigint NOT NULL, "mission_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- set comment to column: "started_at" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."started_at" IS '任务开始时刻';
-- set comment to column: "finished_at" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."finished_at" IS '任务完成时刻';
-- set comment to column: "result" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."result" IS '任务结果';
-- set comment to column: "device_id" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."device_id" IS '领到任务的设备 ID';
-- set comment to column: "result_urls" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."result_urls" IS '任务结果资源位置列表序列化';
-- set comment to column: "key_pair_id" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."key_pair_id" IS '密钥对 ID';
-- set comment to column: "mission_id" on table: "mission_key_pairs"
COMMENT ON COLUMN "mission_key_pairs" ."mission_id" IS '任务 ID';
-- create "mission_kinds" table
CREATE TABLE "mission_kinds" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'txt2img', "category" character varying NOT NULL DEFAULT 'SD', "billing_type" character varying NOT NULL DEFAULT 'count', PRIMARY KEY ("id"));
-- set comment to column: "type" on table: "mission_kinds"
COMMENT ON COLUMN "mission_kinds" ."type" IS '任务单位类型';
-- set comment to column: "category" on table: "mission_kinds"
COMMENT ON COLUMN "mission_kinds" ."category" IS '任务大类';
-- set comment to column: "billing_type" on table: "mission_kinds"
COMMENT ON COLUMN "mission_kinds" ."billing_type" IS '计费类型';
-- create "mission_produce_orders" table
CREATE TABLE "mission_produce_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "mission_id" bigint NOT NULL DEFAULT 0, "status" character varying NOT NULL DEFAULT 'waiting', "pure_cep" bigint NOT NULL DEFAULT 0, "gift_cep" bigint NOT NULL DEFAULT 0, "type" character varying NOT NULL DEFAULT 'txt2img', "is_time" boolean NOT NULL DEFAULT false, "serial_number" character varying NOT NULL DEFAULT '', "device_id" bigint NOT NULL DEFAULT 0, "mission_consume_order_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "mission_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_id" IS '任务 id，关联任务中枢的任务';
-- set comment to column: "status" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."status" IS '任务订单的状态，注意不强关联任务的状态';
-- set comment to column: "pure_cep" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."pure_cep" IS '任务收益的本金 cep 量';
-- set comment to column: "gift_cep" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."gift_cep" IS '任务收益的赠送 cep 量';
-- set comment to column: "type" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."type" IS '任务类型，计时或者次数任务';
-- set comment to column: "is_time" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."is_time" IS '是否为计时类型任务';
-- set comment to column: "serial_number" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."serial_number" IS '订单序列号';
-- set comment to column: "device_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."device_id" IS '生产者接该任务用的设备 id';
-- set comment to column: "mission_consume_order_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_consume_order_id" IS '外键任务消费订单，一个任务消费订单可能会对应多个任务生产订单';
-- set comment to column: "user_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."user_id" IS '外键关联用户 id';
-- create "output_logs" table
CREATE TABLE "output_logs" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "trace_id" bigint NOT NULL, "headers" character varying NOT NULL, "body" character varying NULL DEFAULT '', "url" character varying NOT NULL, "ip" character varying NULL DEFAULT '', "caller" character varying NOT NULL DEFAULT 'general', "status" smallint NULL DEFAULT 200, "hmac_key" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- set comment to column: "trace_id" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."trace_id" IS '请求追踪 id';
-- set comment to column: "headers" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."headers" IS '请求头';
-- set comment to column: "body" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."body" IS '请求体';
-- set comment to column: "url" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."url" IS '请求地址';
-- set comment to column: "ip" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."ip" IS '客户端 IP';
-- set comment to column: "caller" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."caller" IS '调用方';
-- set comment to column: "status" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."status" IS '状态码';
-- set comment to column: "hmac_key" on table: "output_logs"
COMMENT ON COLUMN "output_logs" ."hmac_key" IS '记录调用者的密钥对';
-- create "platform_accounts" table
CREATE TABLE "platform_accounts" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'profit', "sum_total_cep" bigint NOT NULL DEFAULT 0, "total_cep" bigint NOT NULL DEFAULT 0, "sum_pure_cep" bigint NOT NULL DEFAULT 0, "pure_cep" bigint NOT NULL DEFAULT 0, "sum_gift_cep" bigint NOT NULL DEFAULT 0, "gift_cep" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "platformaccount_type_deleted_at" to table: "platform_accounts"
CREATE UNIQUE INDEX "platformaccount_type_deleted_at" ON "platform_accounts" ("type", "deleted_at");
-- set comment to column: "sum_total_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."sum_total_cep" IS '累计总余额';
-- set comment to column: "total_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."total_cep" IS '剩余总余额';
-- set comment to column: "sum_pure_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."sum_pure_cep" IS '累计本金';
-- set comment to column: "pure_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."pure_cep" IS '剩余本金';
-- set comment to column: "sum_gift_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."sum_gift_cep" IS '累计赠金';
-- set comment to column: "gift_cep" on table: "platform_accounts"
COMMENT ON COLUMN "platform_accounts" ."gift_cep" IS '剩余赠金';
-- create "profit_accounts" table
CREATE TABLE "profit_accounts" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "sum_cep" bigint NOT NULL DEFAULT 0, "remain_cep" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "profit_accounts_user_id_key" to table: "profit_accounts"
CREATE UNIQUE INDEX "profit_accounts_user_id_key" ON "profit_accounts" ("user_id");
-- set comment to column: "sum_cep" on table: "profit_accounts"
COMMENT ON COLUMN "profit_accounts" ."sum_cep" IS '累计分润余额';
-- set comment to column: "remain_cep" on table: "profit_accounts"
COMMENT ON COLUMN "profit_accounts" ."remain_cep" IS '剩余分润余额，未提现的';
-- set comment to column: "user_id" on table: "profit_accounts"
COMMENT ON COLUMN "profit_accounts" ."user_id" IS '外键用户 id';
-- create "profit_settings" table
CREATE TABLE "profit_settings" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "ratio" bigint NOT NULL DEFAULT 75, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "ratio" on table: "profit_settings"
COMMENT ON COLUMN "profit_settings" ."ratio" IS '分润比例';
-- set comment to column: "user_id" on table: "profit_settings"
COMMENT ON COLUMN "profit_settings" ."user_id" IS '外键用户 id';
-- create "recharge_orders" table
CREATE TABLE "recharge_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'pending', "pure_cep" bigint NOT NULL DEFAULT 0, "type" character varying NOT NULL DEFAULT 'vx', "serial_number" character varying NOT NULL DEFAULT '', "third_api_resp" character varying NOT NULL DEFAULT '', "from_user_id" bigint NOT NULL DEFAULT 0, "out_transaction_id" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, "social_id" bigint NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "status" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."status" IS '充值订单的状态，比如微信发起支付后可能没完成支付';
-- set comment to column: "pure_cep" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."pure_cep" IS '充值多少本金';
-- set comment to column: "type" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."type" IS '充值订单的类型';
-- set comment to column: "serial_number" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."serial_number" IS '充值订单的序列号';
-- set comment to column: "third_api_resp" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."third_api_resp" IS '第三方平台的返回，给到前端才能发起支付';
-- set comment to column: "from_user_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."from_user_id" IS '由谁发起的充值';
-- set comment to column: "out_transaction_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."out_transaction_id" IS '平台方订单号';
-- set comment to column: "user_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."user_id" IS '充值的用户 id';
-- set comment to column: "social_id" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."social_id" IS '关联充值来源的身份源 id';
-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '请设置昵称', "jpg_url" character varying NOT NULL DEFAULT '', "key" character varying NOT NULL DEFAULT '', "secret" character varying NOT NULL DEFAULT '', "phone" character varying NOT NULL DEFAULT '', "password" character varying NOT NULL DEFAULT '', "is_frozen" boolean NOT NULL DEFAULT false, "user_type" character varying NOT NULL DEFAULT 'personal', "parent_id" bigint NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "name" on table: "users"
COMMENT ON COLUMN "users" ."name" IS '用户名';
-- set comment to column: "jpg_url" on table: "users"
COMMENT ON COLUMN "users" ."jpg_url" IS '头像';
-- set comment to column: "key" on table: "users"
COMMENT ON COLUMN "users" ."key" IS '用户在任务中枢密钥对的键';
-- set comment to column: "secret" on table: "users"
COMMENT ON COLUMN "users" ."secret" IS '用户在任务中枢的密钥对的值';
-- set comment to column: "phone" on table: "users"
COMMENT ON COLUMN "users" ."phone" IS '用户的手机号';
-- set comment to column: "password" on table: "users"
COMMENT ON COLUMN "users" ."password" IS '密码';
-- set comment to column: "is_frozen" on table: "users"
COMMENT ON COLUMN "users" ."is_frozen" IS '是否冻结';
-- set comment to column: "user_type" on table: "users"
COMMENT ON COLUMN "users" ."user_type" IS '用户类型';
-- set comment to column: "parent_id" on table: "users"
COMMENT ON COLUMN "users" ."parent_id" IS '邀请人用户 id';
-- create "user_devices" table
CREATE TABLE "user_devices" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "device_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "device_id" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."device_id" IS '外键设备 id';
-- set comment to column: "user_id" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."user_id" IS '外键用户 id';
-- create "vx_accounts" table
CREATE TABLE "vx_accounts" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "open_id" character varying NOT NULL DEFAULT '', "union_id" character varying NOT NULL DEFAULT '', "scope" character varying NOT NULL DEFAULT 'base', "session_key" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "open_id" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."open_id" IS '微信账户的 open_id';
-- set comment to column: "union_id" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."union_id" IS '微信账户的 union_id';
-- set comment to column: "scope" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."scope" IS '账户的权限级别，一般为 base';
-- set comment to column: "session_key" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."session_key" IS '会话密钥';
-- set comment to column: "user_id" on table: "vx_accounts"
COMMENT ON COLUMN "vx_accounts" ."user_id" IS '外键用户 id';
-- create "vx_socials" table
CREATE TABLE "vx_socials" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "app_id" character varying NOT NULL DEFAULT '', "open_id" character varying NOT NULL DEFAULT '', "union_id" character varying NOT NULL DEFAULT '', "scope" character varying NOT NULL DEFAULT 'base', "session_key" character varying NOT NULL DEFAULT '', "access_token" character varying NOT NULL DEFAULT '', "refresh_token" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "app_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."app_id" IS '微信应用 id';
-- set comment to column: "open_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."open_id" IS '微信身份源的 open_id';
-- set comment to column: "union_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."union_id" IS '微信身份源的 union_id';
-- set comment to column: "scope" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."scope" IS '账户的权限级别，一般为 base';
-- set comment to column: "session_key" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."session_key" IS '小程序专用的会话密钥，不可以返回给前端';
-- set comment to column: "access_token" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."access_token" IS '微信能力访问凭证';
-- set comment to column: "refresh_token" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."refresh_token" IS '刷新微信凭证的刷新凭证';
-- set comment to column: "user_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."user_id" IS '外键用户 id';
