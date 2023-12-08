-- create "withdraw_accounts" table
CREATE TABLE "withdraw_accounts" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "business_name" character varying NOT NULL DEFAULT '0', "business_type" character varying NOT NULL DEFAULT '0', "business_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "withdraw_accounts"
COMMENT ON TABLE "withdraw_accounts" IS '提现账户，用来提供提现渠道';
-- set comment to column: "id" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "business_name" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_name" IS '商户名称';
-- set comment to column: "business_type" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_type" IS '商户名称';
-- set comment to column: "business_id" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."business_id" IS '货币余额';
-- set comment to column: "user_id" on table: "withdraw_accounts"
COMMENT ON COLUMN "withdraw_accounts" ."user_id" IS '外键用户 id';
