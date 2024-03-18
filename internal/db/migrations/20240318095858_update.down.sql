-- reverse: set comment to column: "domain" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."domain" IS '';
-- reverse: modify "frps_infos" table
ALTER TABLE "frps_infos" DROP COLUMN "domain";
