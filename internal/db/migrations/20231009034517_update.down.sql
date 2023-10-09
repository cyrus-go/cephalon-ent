-- reverse: set comment to column: "renewal_type" on table: "prices"
COMMENT ON COLUMN "prices" ."renewal_type" IS '';
-- reverse: modify "prices" table
ALTER TABLE "prices" DROP COLUMN "renewal_type";
