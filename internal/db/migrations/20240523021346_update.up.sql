-- create "withdraw_records" table
CREATE TABLE "withdraw_records" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "withdraw_account" character varying NOT NULL DEFAULT '', "type" character varying NOT NULL DEFAULT 'unknown', "amount" bigint NOT NULL DEFAULT 0, "remain_amount" bigint NOT NULL DEFAULT 0, "rate" bigint NOT NULL DEFAULT 7, "real_amount" bigint NOT NULL DEFAULT 0, "status" character varying NOT NULL DEFAULT 'pending', "reject_reason" character varying NOT NULL DEFAULT '', "transfer_order_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, "operate_user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "withdraw_records_transfer_order_id_key" to table: "withdraw_records"
CREATE UNIQUE INDEX "withdraw_records_transfer_order_id_key" ON "withdraw_records" ("transfer_order_id");
-- create index "withdrawrecord_user_id" to table: "withdraw_records"
CREATE INDEX "withdrawrecord_user_id" ON "withdraw_records" ("user_id");
-- set comment to table: "withdraw_records"
COMMENT ON TABLE "withdraw_records" IS '提现记录，记录所有的提现信息';
-- set comment to column: "id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "withdraw_account" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."withdraw_account" IS '提现账户';
-- set comment to column: "type" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."type" IS '提现类型';
-- set comment to column: "amount" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."amount" IS '提现金额，单位：厘';
-- set comment to column: "remain_amount" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."remain_amount" IS '本次提现后余额，单位：厘';
-- set comment to column: "rate" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."rate" IS '提现手续费率，100 为基准，比如手续费 7%，值就应该为 7，最大值不能超过 100, 默认 7%';
-- set comment to column: "real_amount" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."real_amount" IS '提现实际到账（扣除手续费），单位：厘';
-- set comment to column: "status" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."status" IS '转账订单的状态，比如微信发起支付后可能没完成支付';
-- set comment to column: "reject_reason" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."reject_reason" IS '提现审批拒绝的理由';
-- set comment to column: "transfer_order_id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."transfer_order_id" IS '对应的交易订单 id（一对一）';
-- set comment to column: "user_id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."user_id" IS '提现的用户 id';
-- set comment to column: "operate_user_id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."operate_user_id" IS '操作的用户 id，手动充值或者提现审批才有数据，默认为 0';
