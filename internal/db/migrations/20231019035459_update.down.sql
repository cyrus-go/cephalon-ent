-- reverse: set comment to column: "gpu_id" on table: "prices"
COMMENT ON COLUMN "prices" ."gpu_id" IS '';
-- reverse: modify "prices" table
ALTER TABLE "prices" DROP COLUMN "gpu_id";
-- reverse: set comment to column: "memory" on table: "gpus"
COMMENT ON COLUMN "gpus" ."memory" IS '';
-- reverse: set comment to column: "cpu" on table: "gpus"
COMMENT ON COLUMN "gpus" ."cpu" IS '';
-- reverse: set comment to column: "video_memory" on table: "gpus"
COMMENT ON COLUMN "gpus" ."video_memory" IS '';
-- reverse: modify "gpus" table
ALTER TABLE "gpus" DROP COLUMN "memory", DROP COLUMN "cpu", DROP COLUMN "video_memory";
