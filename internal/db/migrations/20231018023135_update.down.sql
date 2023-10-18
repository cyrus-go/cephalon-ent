-- reverse: set comment to column: "device_id" on table: "missions"
COMMENT ON COLUMN "missions" ."device_id" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "device_id";
