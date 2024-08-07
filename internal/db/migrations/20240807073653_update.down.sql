-- reverse: set comment to column: "rank" on table: "devices"
COMMENT ON COLUMN "devices" ."rank" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "rank";
