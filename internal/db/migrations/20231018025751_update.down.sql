-- reverse: set comment to column: "black_device_ids" on table: "missions"
COMMENT ON COLUMN "missions" ."black_device_ids" IS '';
-- reverse: set comment to column: "white_device_ids" on table: "missions"
COMMENT ON COLUMN "missions" ."white_device_ids" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "black_device_ids", DROP COLUMN "white_device_ids";
