-- modify "cdk_infos" table
ALTER TABLE "cdk_infos" ADD COLUMN "status" character varying NOT NULL DEFAULT 'unknown';
-- set comment to column: "status" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."status" IS 'cdk 状态';
