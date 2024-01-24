-- reverse: set comment to column: "status" on table: "cdk_infos"
COMMENT ON COLUMN "cdk_infos" ."status" IS '';
-- reverse: modify "cdk_infos" table
ALTER TABLE "cdk_infos" DROP COLUMN "status";
