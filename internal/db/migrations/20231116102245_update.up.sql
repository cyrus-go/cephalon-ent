-- modify "prices" table
ALTER TABLE "prices" ADD COLUMN "is_hot_gpu" boolean NOT NULL DEFAULT false;
-- set comment to column: "is_hot_gpu" on table: "prices"
COMMENT ON COLUMN "prices" ."is_hot_gpu" IS '是否为热点 Gpu';
