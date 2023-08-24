-- reverse: set comment to column: "gpu_version" on table: "missions"
COMMENT ON COLUMN "missions" ."gpu_version" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "gpu_version";
