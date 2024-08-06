-- reverse: set comment to column: "timed_shutdown" on table: "missions"
COMMENT ON COLUMN "missions" ."timed_shutdown" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "timed_shutdown";
