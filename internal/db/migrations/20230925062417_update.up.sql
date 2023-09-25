-- modify "bills" table
ALTER TABLE "bills" ADD COLUMN "mission_order_bills" bigint NULL;
-- create "mission_orders" table
CREATE TABLE "mission_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'unknown', "consume_amount" bigint NOT NULL DEFAULT 0, "produce_amount" bigint NOT NULL DEFAULT 0, "gas_amount" bigint NOT NULL DEFAULT 0, "mission_type" character varying NOT NULL DEFAULT 'unknown', "mission_billing_type" character varying NOT NULL DEFAULT 'unknown', "call_way" character varying NOT NULL DEFAULT 'unknown', "serial_number" character varying NOT NULL DEFAULT '', "started_at" timestamptz NOT NULL, "finished_at" timestamptz NOT NULL, "mission_batch_number" character varying NOT NULL DEFAULT '', "mission_id" bigint NOT NULL DEFAULT 0, "mission_batch_id" bigint NOT NULL DEFAULT 0, "symbol_id" bigint NOT NULL DEFAULT 0, "consume_user_id" bigint NOT NULL DEFAULT 0, "produce_user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "mission_orders"
COMMENT ON TABLE "mission_orders" IS '任务订单，记录由任务产生的一些金额变动情况，取代任务消耗订单和任务生产订单';
-- set comment to column: "status" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."status" IS '任务订单的状态，注意不强关联任务的状态';
-- set comment to column: "consume_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."consume_amount" IS '订单的货币消耗量';
-- set comment to column: "produce_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."produce_amount" IS '订单的货币分润量';
-- set comment to column: "gas_amount" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."gas_amount" IS '订单的平台收量，不记录用户 id，因为都是记载到 genesis 用户';
-- set comment to column: "mission_type" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_type" IS '任务类型，等于任务表的类型字段';
-- set comment to column: "mission_billing_type" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_billing_type" IS '是否为计时类型任务';
-- set comment to column: "call_way" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."call_way" IS '调用方式，API 调用或者微信小程序调用';
-- set comment to column: "serial_number" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."serial_number" IS '订单序列号';
-- set comment to column: "started_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."started_at" IS '任务开始执行时刻';
-- set comment to column: "finished_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."finished_at" IS '任务结束执行时刻';
-- set comment to column: "mission_batch_number" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_batch_number" IS '任务批次号，用于方便检索';
-- set comment to column: "mission_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_id" IS '任务 id，外键关联任务';
-- set comment to column: "mission_batch_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."mission_batch_id" IS '任务批次外键';
-- set comment to column: "symbol_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."symbol_id" IS '币种 id';
-- set comment to column: "consume_user_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."consume_user_id" IS '外键关联发起任务的用户 id';
-- set comment to column: "produce_user_id" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."produce_user_id" IS '外键关联完成任务的用户 id';
