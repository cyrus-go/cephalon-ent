-- reverse: set comment to column: "sd_api" on table: "missions"
COMMENT ON COLUMN "missions" ."sd_api" IS '';
-- reverse: set comment to column: "resp_body" on table: "missions"
COMMENT ON COLUMN "missions" ."resp_body" IS '';
-- reverse: set comment to column: "resp_status_code" on table: "missions"
COMMENT ON COLUMN "missions" ."resp_status_code" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "sd_api", DROP COLUMN "resp_body", DROP COLUMN "resp_status_code";
