-- reverse: set comment to column: "is_hot_gpu" on table: "prices"
COMMENT ON COLUMN "prices" ."is_hot_gpu" IS '';
-- reverse: modify "prices" table
ALTER TABLE "prices" DROP COLUMN "is_hot_gpu";
