-- create "cdk_infos" table
CREATE TABLE "cdk_infos" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "cdk_number" character varying NOT NULL DEFAULT '', "type" character varying NOT NULL DEFAULT 'unknown', "get_cep" bigint NOT NULL DEFAULT 0, "get_time" bigint NOT NULL DEFAULT 0, "billing_type" character varying NOT NULL DEFAULT 'unknown', "expired_at" timestamptz NULL, "use_times" bigint NOT NULL DEFAULT 0, "issue_user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "cdkinfo_cdk_number" to table: "cdk_infos"
CREATE INDEX "cdkinfo_cdk_number" ON "cdk_infos" ("cdk_number");
-- set comment to table: "cdk_infos"
COMMENT ON TABLE "cdk_infos" IS '兑换码，可以兑换脑力值、gpu 使用权等';
-- set comment to column: "id" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "cdk_number" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."cdk_number" IS 'cdk 序列号';
-- set comment to column: "type" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."type" IS 'cdk 类型';
-- set comment to column: "get_cep" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."get_cep" IS 'cdk 能兑换的 cep 数量';
-- set comment to column: "get_time" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."get_time" IS 'cdk 能兑换的 gpu 使用时长';
-- set comment to column: "billing_type" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."billing_type" IS '兑换 gpu 使用时长的类型';
-- set comment to column: "expired_at" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."expired_at" IS '过期时间';
-- set comment to column: "use_times" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."use_times" IS 'cdk 能使用的次数';
-- set comment to column: "issue_user_id" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."issue_user_id" IS '外键：发行用户 id';
