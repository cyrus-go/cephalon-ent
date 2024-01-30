-- modify "cdk_infos" table
ALTER TABLE "cdk_infos" ADD COLUMN "used_at" timestamptz NULL, ADD COLUMN "use_user_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "used_at" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."used_at" IS '使用时间';
-- set comment to column: "use_user_id" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."use_user_id" IS '外键：使用 cdk 用户 id';
