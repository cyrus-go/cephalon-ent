-- reverse: set comment to column: "group" on table: "surveys"
COMMENT ON COLUMN "surveys" ."group" IS '';
-- reverse: set comment to column: "sort_num" on table: "surveys"
COMMENT ON COLUMN "surveys" ."sort_num" IS '';
-- reverse: set comment to column: "ended_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."ended_at" IS '';
-- reverse: set comment to column: "started_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."started_at" IS '';
-- reverse: modify "surveys" table
ALTER TABLE "surveys" DROP COLUMN "group", DROP COLUMN "sort_num", DROP COLUMN "ended_at", DROP COLUMN "started_at";
