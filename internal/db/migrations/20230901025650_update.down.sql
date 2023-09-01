-- reverse: set comment to column: "is_deprecated" on table: "prices"
COMMENT ON COLUMN "prices" ."is_deprecated" IS '';
-- reverse: modify "prices" table
ALTER TABLE "prices" DROP COLUMN "is_deprecated";
