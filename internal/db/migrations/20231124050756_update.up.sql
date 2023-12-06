-- modify "mission_orders" table
ALTER TABLE "mission_orders" ADD COLUMN "lately_settled_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00.0000000 +00:00';
-- set comment to column: "buy_duration" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."buy_duration" IS '包时任务订单购买的时长';
-- set comment to column: "lately_settled_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."lately_settled_at" IS '上一次结算时间';
-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "free_at" timestamptz default '0001-01-01 00:00:00.00000 +00:00' NOT NULL, ADD COLUMN "extra_service_missions" bigint NULL;
-- set comment to column: "free_at" on table: "missions"
COMMENT ON COLUMN "missions" ."free_at" IS '任务释放时刻';
-- create "extra_services" table
CREATE TABLE "extra_services" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', "extra_service_type" character varying NOT NULL DEFAULT 'unknown', "started_at" timestamptz NULL, "finished_at" timestamptz NULL, "mission_extra_services" bigint NULL, PRIMARY KEY ("id"));
-- set comment to table: "extra_services"
COMMENT ON TABLE "extra_services" IS '附加服务表';
-- set comment to column: "id" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "name" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."name" IS '附加服务名称';
-- set comment to column: "extra_service_type" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."extra_service_type" IS '附加服务类型';
-- set comment to column: "started_at" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."started_at" IS '附加服务购买时间开始，为空表示永久有效';
-- set comment to column: "finished_at" on table: "extra_services"
COMMENT ON COLUMN "extra_services" ."finished_at" IS '附加服务购买时间结束，为空表示永久有效';
-- create "extra_service_orders" table
CREATE TABLE "extra_service_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "amount" bigint NOT NULL DEFAULT 0, "extra_service_type" character varying NOT NULL DEFAULT 'unknown', "buy_duration" bigint NOT NULL DEFAULT 0, "plan_started_at" timestamptz NULL, "plan_finished_at" timestamptz NULL, "mission_id" bigint NOT NULL DEFAULT 0, "mission_batch_id" bigint NOT NULL DEFAULT 0, "symbol_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "extra_service_orders"
COMMENT ON TABLE "extra_service_orders" IS '附加服务订单表';
-- set comment to column: "id" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "amount" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."amount" IS '订单的货币消耗量';
-- set comment to column: "extra_service_type" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."extra_service_type" IS '附加服务类型';
-- set comment to column: "buy_duration" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."buy_duration" IS '包时任务订单购买的时长';
-- set comment to column: "plan_started_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."plan_started_at" IS '任务计划开始时间（包时）';
-- set comment to column: "plan_finished_at" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."plan_finished_at" IS '任务计划结束时间（包时）';
-- set comment to column: "mission_id" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."mission_id" IS '任务 id，外键关联任务';
-- set comment to column: "mission_batch_id" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."mission_batch_id" IS '任务批次外键';
-- set comment to column: "symbol_id" on table: "extra_service_orders"
COMMENT ON COLUMN "extra_service_orders" ."symbol_id" IS '币种 id';
-- create "extra_service_prices" table
CREATE TABLE "extra_service_prices" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "extra_service_billing_type" character varying NOT NULL DEFAULT 'unknown', "cep" bigint NOT NULL DEFAULT 0, "started_at" timestamptz NULL, "finished_at" timestamptz NULL, "is_deprecated" boolean NOT NULL DEFAULT false, "is_sensitive" boolean NOT NULL DEFAULT false, "extra_service_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "extra_service_prices"
COMMENT ON TABLE "extra_service_prices" IS '附加服务价格表';
-- set comment to column: "id" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "extra_service_billing_type" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."extra_service_billing_type" IS '附加服务订单类型';
-- set comment to column: "cep" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."cep" IS '附加服务单价';
-- set comment to column: "started_at" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."started_at" IS '价格有效时间开始，为空表示永久有效';
-- set comment to column: "finished_at" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."finished_at" IS '价格有效时间结束，为空表示永久有效';
-- set comment to column: "is_deprecated" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."is_deprecated" IS '价格是否屏蔽，前端置灰，硬选也可以';
-- set comment to column: "is_sensitive" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."is_sensitive" IS '价格是否敏感，用于特殊类型任务，不能让外部看到选项';
-- set comment to column: "extra_service_id" on table: "extra_service_prices"
COMMENT ON COLUMN "extra_service_prices" ."extra_service_id" IS '附加服务 id，外键关联附加服务';
-- create "mission_extra_services" table
CREATE TABLE "mission_extra_services" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "extra_service_id" bigint NOT NULL DEFAULT 0, "mission_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "mission_extra_services"
COMMENT ON TABLE "mission_extra_services" IS '任务与附加服务的中间关系';
-- set comment to column: "id" on table: "mission_extra_services"
COMMENT ON COLUMN "mission_extra_services" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "mission_extra_services"
COMMENT ON COLUMN "mission_extra_services" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "mission_extra_services"
COMMENT ON COLUMN "mission_extra_services" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "mission_extra_services"
COMMENT ON COLUMN "mission_extra_services" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "mission_extra_services"
COMMENT ON COLUMN "mission_extra_services" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "mission_extra_services"
COMMENT ON COLUMN "mission_extra_services" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "extra_service_id" on table: "mission_extra_services"
COMMENT ON COLUMN "mission_extra_services" ."extra_service_id" IS '外键附加服务 id';
-- set comment to column: "mission_id" on table: "mission_extra_services"
COMMENT ON COLUMN "mission_extra_services" ."mission_id" IS '外键任务 id';
