-- reverse: set comment to column: "expired_at" on table: "missions"
COMMENT ON COLUMN "missions" ."expired_at" IS '';
-- reverse: set comment to column: "started_at" on table: "missions"
COMMENT ON COLUMN "missions" ."started_at" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "expired_at", DROP COLUMN "started_at";
