-- reverse: set comment to column: "gpu_num" on table: "missions"
COMMENT ON COLUMN "missions" ."gpu_num" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "gpu_num";
