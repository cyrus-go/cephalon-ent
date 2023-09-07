-- reverse: set comment to column: "is_using" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."is_using" IS '';
-- reverse: modify "frpc_infos" table
ALTER TABLE "frpc_infos" DROP COLUMN "is_using";
