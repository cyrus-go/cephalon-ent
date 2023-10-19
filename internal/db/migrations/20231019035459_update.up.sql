-- modify "gpus" table
ALTER TABLE "gpus" ADD COLUMN "video_memory" bigint NOT NULL DEFAULT 40, ADD COLUMN "cpu" bigint NOT NULL DEFAULT 12, ADD COLUMN "memory" bigint NOT NULL DEFAULT 128;
-- set comment to column: "video_memory" on table: "gpus"
COMMENT ON COLUMN "gpus" ."video_memory" IS '显存';
-- set comment to column: "cpu" on table: "gpus"
COMMENT ON COLUMN "gpus" ."cpu" IS 'CPU';
-- set comment to column: "memory" on table: "gpus"
COMMENT ON COLUMN "gpus" ."memory" IS '内存';
-- modify "prices" table
ALTER TABLE "prices" ADD COLUMN "gpu_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "gpu_id" on table: "prices"
COMMENT ON COLUMN "prices" ."gpu_id" IS '外键 gpu id';
