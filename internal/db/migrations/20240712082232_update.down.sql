-- reverse: set comment to column: "remark" on table: "missions"
COMMENT ON COLUMN "missions" ."remark" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "remark";
