-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "free_gpu_num" bigint NOT NULL DEFAULT 1;
-- set comment to column: "free_gpu_num" on table: "devices"
COMMENT ON COLUMN "devices" ."free_gpu_num" IS '空闲GPU数量';
