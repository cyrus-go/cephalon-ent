-- reverse: set comment to column: "closed_at" on table: "missions"
COMMENT ON COLUMN "missions" ."closed_at" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "closed_at";
