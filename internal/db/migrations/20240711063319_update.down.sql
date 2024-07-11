-- reverse: set comment to column: "hint" on table: "surveys"
COMMENT ON COLUMN "surveys" ."hint" IS '';
-- reverse: modify "surveys" table
ALTER TABLE "surveys" DROP COLUMN "hint";
