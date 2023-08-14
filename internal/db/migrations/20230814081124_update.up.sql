-- set comment to table: "users"
COMMENT ON TABLE "users" IS '用户表，手机号唯一，用户名唯一，hmac_key 唯一';
-- create "bills" table
CREATE TABLE "bills" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'mission_consume', "is_add" boolean NOT NULL DEFAULT false, "serial_number" character varying NOT NULL DEFAULT '', "cep" bigint NOT NULL DEFAULT 0, "status" character varying NOT NULL DEFAULT 'pending', "market_bill_id" bigint NOT NULL DEFAULT 0, "platform_cep" bigint NOT NULL DEFAULT 0, "reason_id" bigint NULL DEFAULT 0, "platform_wallet_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, "wallet_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "bills"
COMMENT ON TABLE "bills" IS '记录额度账户的变动的主流水账单记录';
-- set comment to column: "type" on table: "bills"
COMMENT ON COLUMN "bills" ."type" IS '次级账单流水的类型，充值或者任务消耗或任务收益';
-- set comment to column: "is_add" on table: "bills"
COMMENT ON COLUMN "bills" ."is_add" IS '是否增加余额，布尔值默认为否';
-- set comment to column: "serial_number" on table: "bills"
COMMENT ON COLUMN "bills" ."serial_number" IS '账单序列号';
-- set comment to column: "cep" on table: "bills"
COMMENT ON COLUMN "bills" ."cep" IS '消费或收益多少 cep';
-- set comment to column: "status" on table: "bills"
COMMENT ON COLUMN "bills" ."status" IS '消耗流水一开始不能直接生效，确定关联的消耗时间成功后才可以扣费';
-- set comment to column: "market_bill_id" on table: "bills"
COMMENT ON COLUMN "bills" ."market_bill_id" IS '营销流水 id';
-- set comment to column: "platform_cep" on table: "bills"
COMMENT ON COLUMN "bills" ."platform_cep" IS '平台收取多少本金余额';
-- set comment to column: "reason_id" on table: "bills"
COMMENT ON COLUMN "bills" ."reason_id" IS '关联消耗产生的来源外键 id，比如 type 为 mission 时关联任务订单';
-- set comment to column: "platform_wallet_id" on table: "bills"
COMMENT ON COLUMN "bills" ."platform_wallet_id" IS '平台分润钱包 id';
-- set comment to column: "user_id" on table: "bills"
COMMENT ON COLUMN "bills" ."user_id" IS '外键用户 id';
-- set comment to column: "wallet_id" on table: "bills"
COMMENT ON COLUMN "bills" ."wallet_id" IS '外键钱包 id';
-- create "collections" table
CREATE TABLE "collections" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "url" character varying NOT NULL DEFAULT '', "picture_name" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "collections"
COMMENT ON TABLE "collections" IS '资源收藏，与用户一对多';
-- set comment to column: "url" on table: "collections"
COMMENT ON COLUMN "collections" ."url" IS '路径';
-- set comment to column: "picture_name" on table: "collections"
COMMENT ON COLUMN "collections" ."picture_name" IS '照片名称';
-- set comment to column: "user_id" on table: "collections"
COMMENT ON COLUMN "collections" ."user_id" IS '外键用户 id';
-- create "devices" table
CREATE TABLE "devices" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'online', "binding_status" character varying NOT NULL DEFAULT 'init', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "devices"
COMMENT ON TABLE "devices" IS '设备表，与用户一对多';
-- set comment to column: "status" on table: "devices"
COMMENT ON COLUMN "devices" ."status" IS '设备状态';
-- set comment to column: "binding_status" on table: "devices"
COMMENT ON COLUMN "devices" ."binding_status" IS '设备的绑定状态';
-- set comment to column: "user_id" on table: "devices"
COMMENT ON COLUMN "devices" ."user_id" IS '外键用户 id';
-- create "hmac_key_pairs" table
CREATE TABLE "hmac_key_pairs" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "key" character varying NOT NULL DEFAULT '', "secret" character varying NOT NULL DEFAULT '', "caller" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, "user_hmac_key_pair" bigint NOT NULL, PRIMARY KEY ("id"));
-- create index "hmac_key_pairs_user_hmac_key_pair_key" to table: "hmac_key_pairs"
CREATE UNIQUE INDEX "hmac_key_pairs_user_hmac_key_pair_key" ON "hmac_key_pairs" ("user_hmac_key_pair");
-- create index "hmackeypair_key" to table: "hmac_key_pairs"
CREATE INDEX "hmackeypair_key" ON "hmac_key_pairs" ("key");
-- set comment to table: "hmac_key_pairs"
COMMENT ON TABLE "hmac_key_pairs" IS 'Hmac 密钥对，用于没有登录态时安全调用任务相关接口的场景';
-- set comment to column: "key" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."key" IS '密钥对的 key 值，用于检索密钥';
-- set comment to column: "secret" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."secret" IS '加密密钥';
-- set comment to column: "caller" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."caller" IS '请求方';
-- set comment to column: "user_id" on table: "hmac_key_pairs"
COMMENT ON COLUMN "hmac_key_pairs" ."user_id" IS '外键用户 ID';
-- create "input_logs" table
CREATE TABLE "input_logs" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "trace_id" bigint NOT NULL DEFAULT 0, "headers" character varying NOT NULL DEFAULT '', "body" character varying NULL DEFAULT '', "query" character varying NULL DEFAULT '', "url" character varying NOT NULL DEFAULT '', "ip" character varying NOT NULL DEFAULT '', "caller" character varying NOT NULL DEFAULT 'general', "method" character varying NOT NULL DEFAULT 'GET', "hmac_key" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- set comment to table: "input_logs"
COMMENT ON TABLE "input_logs" IS '接口输入表，用于记录最老的直接任务模式';
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
CREATE TABLE "missions" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'txt2img', "is_time" boolean NOT NULL DEFAULT false, "body" character varying NOT NULL DEFAULT '', "call_back_url" character varying NOT NULL DEFAULT '', "status" character varying NOT NULL DEFAULT 'waiting', "result_urls" character varying NOT NULL DEFAULT '', "additional_result" character varying NOT NULL DEFAULT '', "mission_batch_number" character varying NOT NULL DEFAULT '', "hmac_key_pair_id" bigint NOT NULL DEFAULT 0, "mission_batch_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "type" on table: "missions"
COMMENT ON COLUMN "missions" ."type" IS '任务类型';
-- set comment to column: "is_time" on table: "missions"
COMMENT ON COLUMN "missions" ."is_time" IS '是否为计时类型任务';
-- set comment to column: "body" on table: "missions"
COMMENT ON COLUMN "missions" ."body" IS '任务的请求参数体';
-- set comment to column: "call_back_url" on table: "missions"
COMMENT ON COLUMN "missions" ."call_back_url" IS '回调地址，空字符串表示不回调';
-- set comment to column: "status" on table: "missions"
COMMENT ON COLUMN "missions" ."status" IS '任务状态';
-- set comment to column: "result_urls" on table: "missions"
COMMENT ON COLUMN "missions" ."result_urls" IS '任务结果资源位置列表序列化';
-- set comment to column: "additional_result" on table: "missions"
COMMENT ON COLUMN "missions" ."additional_result" IS '有的任务除了链接外还有其他有用的结果，都塞在这个字段，比如 sd 的实际入参';
-- set comment to column: "mission_batch_number" on table: "missions"
COMMENT ON COLUMN "missions" ."mission_batch_number" IS '任务批次号';
-- set comment to column: "hmac_key_pair_id" on table: "missions"
COMMENT ON COLUMN "missions" ."hmac_key_pair_id" IS '外键任务创建者的密钥对 ID';
-- set comment to column: "mission_batch_id" on table: "missions"
COMMENT ON COLUMN "missions" ."mission_batch_id" IS '外键任务批次';
-- set comment to column: "user_id" on table: "missions"
COMMENT ON COLUMN "missions" ."user_id" IS '外键任务创建者的 ID';
-- create "mission_batches" table
CREATE TABLE "mission_batches" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "number" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "mission_batches"
COMMENT ON TABLE "mission_batches" IS '任务批次，和任务相关的数据一对多';
-- set comment to column: "number" on table: "mission_batches"
COMMENT ON COLUMN "mission_batches" ."number" IS '批次号码';
-- set comment to column: "user_id" on table: "mission_batches"
COMMENT ON COLUMN "mission_batches" ."user_id" IS '创建者';
-- create "mission_consume_orders" table
CREATE TABLE "mission_consume_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "mission_id" bigint NOT NULL DEFAULT 0, "status" character varying NOT NULL DEFAULT 'waiting', "cep" bigint NOT NULL DEFAULT 0, "type" character varying NOT NULL DEFAULT 'txt2img', "is_time" boolean NOT NULL DEFAULT false, "call_way" character varying NOT NULL DEFAULT 'api', "serial_number" character varying NOT NULL DEFAULT '', "started_at" timestamptz NULL, "finished_at" timestamptz NULL, "mission_batch_number" character varying NOT NULL DEFAULT '', "mission_mission_consume_order" bigint NOT NULL, "mission_batch_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "mission_consume_orders_mission_mission_consume_order_key" to table: "mission_consume_orders"
CREATE UNIQUE INDEX "mission_consume_orders_mission_mission_consume_order_key" ON "mission_consume_orders" ("mission_mission_consume_order");
-- set comment to table: "mission_consume_orders"
COMMENT ON TABLE "mission_consume_orders" IS '任务消费订单，和任务一对一';
-- set comment to column: "mission_id" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."mission_id" IS '外键任务 id，关联任务';
-- set comment to column: "status" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."status" IS '任务订单的状态，注意不强关联任务的状态';
-- set comment to column: "cep" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."cep" IS '发布任务需消耗的 cep 量';
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
COMMENT ON COLUMN "mission_consume_orders" ."mission_batch_id" IS '外键任务批次 id';
-- set comment to column: "user_id" on table: "mission_consume_orders"
COMMENT ON COLUMN "mission_consume_orders" ."user_id" IS '外键关联用户 id';
-- create "mission_produce_orders" table
CREATE TABLE "mission_produce_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'waiting', "cep" bigint NOT NULL DEFAULT 0, "type" character varying NOT NULL DEFAULT 'txt2img', "is_time" boolean NOT NULL DEFAULT false, "serial_number" character varying NOT NULL DEFAULT '', "device_id" bigint NOT NULL DEFAULT 0, "mission_id" bigint NOT NULL DEFAULT 0, "mission_batch_id" bigint NOT NULL DEFAULT 0, "mission_consume_order_id" bigint NOT NULL DEFAULT 0, "mission_production_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "mission_produce_orders_mission_production_id_key" to table: "mission_produce_orders"
CREATE UNIQUE INDEX "mission_produce_orders_mission_production_id_key" ON "mission_produce_orders" ("mission_production_id");
-- set comment to table: "mission_produce_orders"
COMMENT ON TABLE "mission_produce_orders" IS '任务生产订单';
-- set comment to column: "status" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."status" IS '任务订单的状态，注意不强关联任务的状态';
-- set comment to column: "cep" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."cep" IS '完成任务收益的 cep 量';
-- set comment to column: "type" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."type" IS '任务类型，计时或者次数任务';
-- set comment to column: "is_time" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."is_time" IS '是否为计时类型任务';
-- set comment to column: "serial_number" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."serial_number" IS '订单序列号';
-- set comment to column: "device_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."device_id" IS '生产者接该任务用的设备 id';
-- set comment to column: "mission_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_id" IS '外键任务 id，关联任务中枢的任务';
-- set comment to column: "mission_batch_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_batch_id" IS '外键任务批次';
-- set comment to column: "mission_consume_order_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_consume_order_id" IS '外键任务消费订单，一个任务消费订单可能会对应多个任务生产订单';
-- set comment to column: "mission_production_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."mission_production_id" IS '外键任务生产情况 id';
-- set comment to column: "user_id" on table: "mission_produce_orders"
COMMENT ON COLUMN "mission_produce_orders" ."user_id" IS '外键关联用户 id';
-- create "mission_productions" table
CREATE TABLE "mission_productions" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "started_at" timestamptz NOT NULL, "finished_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'doing', "result_urls" character varying NOT NULL DEFAULT '', "additional_result" character varying NOT NULL DEFAULT '', "device_id" bigint NOT NULL DEFAULT 0, "hmac_key_pair_id" bigint NOT NULL, "mission_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- set comment to table: "mission_productions"
COMMENT ON TABLE "mission_productions" IS '任务生产记录，任务被人接了就会产生一条记录，这一次任务完成情况就在这，同一个任务可以被多次接';
-- set comment to column: "started_at" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."started_at" IS '任务开始时刻';
-- set comment to column: "finished_at" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."finished_at" IS '任务完成时刻';
-- set comment to column: "status" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."status" IS '任务结果';
-- set comment to column: "result_urls" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."result_urls" IS '任务结果资源位置列表序列化';
-- set comment to column: "additional_result" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."additional_result" IS '额外需要返回的结果数据，格式不定';
-- set comment to column: "device_id" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."device_id" IS '领到任务的设备 ID';
-- set comment to column: "hmac_key_pair_id" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."hmac_key_pair_id" IS '密钥对 ID';
-- set comment to column: "mission_id" on table: "mission_productions"
COMMENT ON COLUMN "mission_productions" ."mission_id" IS '任务 ID';
-- create "mission_types" table
CREATE TABLE "mission_types" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'txt2img', "gpu" character varying NOT NULL DEFAULT '3070', "cep" bigint NOT NULL DEFAULT 0, "is_time" boolean NOT NULL DEFAULT false, "category" character varying NOT NULL DEFAULT 'SD', PRIMARY KEY ("id"));
-- set comment to table: "mission_types"
COMMENT ON TABLE "mission_types" IS '任务类型等信息，给任务定价归类等';
-- set comment to column: "type" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."type" IS '任务类型';
-- set comment to column: "gpu" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."gpu" IS '显卡型号';
-- set comment to column: "cep" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."cep" IS '单价消耗 cep';
-- set comment to column: "is_time" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."is_time" IS '是否计时任务';
-- set comment to column: "category" on table: "mission_types"
COMMENT ON COLUMN "mission_types" ."category" IS '任务种类，SD，Jupyter 等';
-- create "output_logs" table
CREATE TABLE "output_logs" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "trace_id" bigint NOT NULL DEFAULT 0, "headers" character varying NOT NULL DEFAULT '', "body" character varying NULL DEFAULT '', "url" character varying NOT NULL DEFAULT '', "ip" character varying NULL DEFAULT '', "caller" character varying NOT NULL DEFAULT 'general', "status" smallint NULL DEFAULT 200, "hmac_key" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- set comment to table: "output_logs"
COMMENT ON TABLE "output_logs" IS '输出表，用于最老的直连方式任务计费';
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
-- create "platform_wallets" table
CREATE TABLE "platform_wallets" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'profit', "sum_cep" bigint NOT NULL DEFAULT 0, "cep" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "platformwallet_type_deleted_at" to table: "platform_wallets"
CREATE UNIQUE INDEX "platformwallet_type_deleted_at" ON "platform_wallets" ("type", "deleted_at");
-- set comment to table: "platform_wallets"
COMMENT ON TABLE "platform_wallets" IS '平台钱包账户，独立于用户，每种类型的 cep 余额，比如分润获取的都集中到一条数据，即一个钱包';
-- set comment to column: "sum_cep" on table: "platform_wallets"
COMMENT ON COLUMN "platform_wallets" ."sum_cep" IS '累计总余额';
-- set comment to column: "cep" on table: "platform_wallets"
COMMENT ON COLUMN "platform_wallets" ."cep" IS '剩余总余额';
-- create "profit_settings" table
CREATE TABLE "profit_settings" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "ratio" bigint NOT NULL DEFAULT 75, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "profit_settings"
COMMENT ON TABLE "profit_settings" IS '用户的分润配置，与用户一对多，但逻辑上主要是一对一';
-- set comment to column: "ratio" on table: "profit_settings"
COMMENT ON COLUMN "profit_settings" ."ratio" IS '分润比例';
-- set comment to column: "user_id" on table: "profit_settings"
COMMENT ON COLUMN "profit_settings" ."user_id" IS '外键用户 id';
-- create "recharge_orders" table
CREATE TABLE "recharge_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'doing', "cep" bigint NOT NULL DEFAULT 0, "type" character varying NOT NULL DEFAULT 'manual', "serial_number" character varying NOT NULL DEFAULT '', "third_api_resp" character varying NOT NULL DEFAULT '', "from_user_id" bigint NOT NULL DEFAULT 0, "out_transaction_id" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, "social_id" bigint NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "recharge_orders"
COMMENT ON TABLE "recharge_orders" IS '充值订单，微信订单、支付宝订单、手动充值都在一张表';
-- set comment to column: "status" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."status" IS '充值订单的状态，比如微信发起支付后可能没完成支付';
-- set comment to column: "cep" on table: "recharge_orders"
COMMENT ON COLUMN "recharge_orders" ."cep" IS '充值多少 cep';
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
-- create "user_devices" table
CREATE TABLE "user_devices" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "device_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "user_devices"
COMMENT ON TABLE "user_devices" IS '用户和设备的记录，主要记录绑定时长和过程';
-- set comment to column: "device_id" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."device_id" IS '外键设备 id';
-- set comment to column: "user_id" on table: "user_devices"
COMMENT ON COLUMN "user_devices" ."user_id" IS '外键用户 id';
-- create "vx_socials" table
CREATE TABLE "vx_socials" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "app_id" character varying NOT NULL DEFAULT '', "open_id" character varying NOT NULL DEFAULT '', "union_id" character varying NOT NULL DEFAULT '', "scope" character varying NOT NULL DEFAULT 'base', "session_key" character varying NOT NULL DEFAULT '', "access_token" character varying NOT NULL DEFAULT '', "refresh_token" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "vx_socials"
COMMENT ON TABLE "vx_socials" IS '微信身份源，与用户一对多，微信订单依赖于它';
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
-- create "wallets" table
CREATE TABLE "wallets" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "cep" bigint NOT NULL DEFAULT 0, "sum_cep" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "wallets_user_id_key" to table: "wallets"
CREATE UNIQUE INDEX "wallets_user_id_key" ON "wallets" ("user_id");
-- set comment to table: "wallets"
COMMENT ON TABLE "wallets" IS '钱包，与用户一对一';
-- set comment to column: "cep" on table: "wallets"
COMMENT ON COLUMN "wallets" ."cep" IS '当前余额';
-- set comment to column: "sum_cep" on table: "wallets"
COMMENT ON COLUMN "wallets" ."sum_cep" IS '累计余额';
-- set comment to column: "user_id" on table: "wallets"
COMMENT ON COLUMN "wallets" ."user_id" IS '外键用户 id';
