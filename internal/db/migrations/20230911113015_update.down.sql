-- reverse: set comment to column: "type" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."type" IS '';
-- reverse: modify "frps_infos" table
ALTER TABLE "frps_infos" DROP COLUMN "type";
