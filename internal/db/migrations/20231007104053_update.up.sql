-- create "renewal_agreements" table
CREATE TABLE "renewal_agreements" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "next_pay_time" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'unknow', "sub_status" character varying NOT NULL DEFAULT 'unknow', "pay_status" character varying NOT NULL DEFAULT 'unknow', "first_pay" bigint NOT NULL DEFAULT 0, "after_pay" bigint NOT NULL DEFAULT 0, "sub_finished_time" timestamptz NOT NULL, "mission_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "renewal_agreements"
COMMENT ON TABLE "renewal_agreements" IS '自动续费协议';
-- set comment to column: "next_pay_time" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."next_pay_time" IS '下次扣款时间';
-- set comment to column: "type" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."type" IS '自动续费类型（按小时、按天等）';
-- set comment to column: "sub_status" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."sub_status" IS '订阅自动续费状态（订阅中、已结束等）';
-- set comment to column: "pay_status" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."pay_status" IS '支付状态（待支付、已支付、支付失败等）';
-- set comment to column: "first_pay" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."first_pay" IS '首次扣款价格';
-- set comment to column: "after_pay" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."after_pay" IS '后续扣款价格';
-- set comment to column: "sub_finished_time" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."sub_finished_time" IS '订阅自动续费结束时间';
-- set comment to column: "mission_id" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."mission_id" IS '外键任务 id';
-- set comment to column: "user_id" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."user_id" IS '外键用户 id';
