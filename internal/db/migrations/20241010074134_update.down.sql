-- reverse: set comment to column: "create_way" on table: "missions"
COMMENT ON COLUMN "missions" ."create_way" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "create_way";
