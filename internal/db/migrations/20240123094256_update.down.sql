-- reverse: set comment to column: "warning_times" on table: "missions"
COMMENT ON COLUMN "missions" ."warning_times" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "warning_times";
