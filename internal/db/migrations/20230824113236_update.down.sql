-- reverse: set comment to column: "power" on table: "gpus"
COMMENT ON COLUMN "gpus" ."power" IS '';
-- reverse: modify "gpus" table
ALTER TABLE "gpus" DROP COLUMN "power";
