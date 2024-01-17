-- reverse: set comment to column: "close_way" on table: "missions"
COMMENT ON COLUMN "missions" ."close_way" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "close_way";
