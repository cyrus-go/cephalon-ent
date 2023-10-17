-- reverse: set comment to column: "password" on table: "missions"
COMMENT ON COLUMN "missions" ."password" IS '';
-- reverse: set comment to column: "username" on table: "missions"
COMMENT ON COLUMN "missions" ."username" IS '';
-- reverse: set comment to column: "call_back_data" on table: "missions"
COMMENT ON COLUMN "missions" ."call_back_data" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "password", DROP COLUMN "username", DROP COLUMN "call_back_data";
