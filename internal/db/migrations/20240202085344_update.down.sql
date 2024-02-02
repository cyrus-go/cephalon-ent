-- reverse: set comment to column: "cep_amount" on table: "prizes"
COMMENT ON COLUMN "prizes" ."cep_amount" IS '';
-- reverse: set comment to column: "type" on table: "prizes"
COMMENT ON COLUMN "prizes" ."type" IS '';
-- reverse: modify "prizes" table
ALTER TABLE "prizes" DROP COLUMN "cep_amount", DROP COLUMN "type";
