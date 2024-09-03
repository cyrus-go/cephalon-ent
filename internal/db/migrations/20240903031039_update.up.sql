-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "gpu_num" bigint NOT NULL DEFAULT 1;
-- set comment to column: "gpu_num" on table: "missions"
COMMENT ON COLUMN "missions" ."gpu_num" IS '多GPU支持,使用GPU数量';
