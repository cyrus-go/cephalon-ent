-- reverse: set comment to column: "use_user_id" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."use_user_id" IS '';
-- reverse: set comment to column: "used_at" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."used_at" IS '';
-- reverse: modify "cdk_infos" table
ALTER TABLE "cdk_infos" DROP COLUMN "use_user_id", DROP COLUMN "used_at";
