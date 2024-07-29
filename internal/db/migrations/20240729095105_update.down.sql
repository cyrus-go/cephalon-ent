-- reverse: set comment to column: "version" on table: "devices"
COMMENT ON COLUMN "devices" ."version" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "version";
