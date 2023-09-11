-- reverse: set comment to column: "token" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."token" IS '';
-- reverse: set comment to column: "authentication_method" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."authentication_method" IS '';
-- reverse: modify "frps_infos" table
ALTER TABLE "frps_infos" DROP COLUMN "token", DROP COLUMN "authentication_method";
