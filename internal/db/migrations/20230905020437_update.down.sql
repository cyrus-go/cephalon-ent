-- reverse: set comment to column: "is_sensitive" on table: "prices"
COMMENT ON COLUMN "prices" ."is_sensitive" IS '';
-- reverse: modify "prices" table
ALTER TABLE "prices" DROP COLUMN "is_sensitive";
