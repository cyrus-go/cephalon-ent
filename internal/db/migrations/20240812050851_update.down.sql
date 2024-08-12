-- reverse: set comment to column: "free_gpu_num" on table: "devices"
COMMENT ON COLUMN "devices" ."free_gpu_num" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "free_gpu_num";
