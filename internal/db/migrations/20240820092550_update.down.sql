-- reverse: set comment to column: "rank_at" on table: "devices"
COMMENT ON COLUMN "devices" ."rank_at" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "rank_at";
