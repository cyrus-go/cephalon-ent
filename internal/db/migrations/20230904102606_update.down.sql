-- reverse: set comment to column: "inner_uri" on table: "missions"
COMMENT ON COLUMN "missions" ."inner_uri" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "inner_uri";
