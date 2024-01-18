-- reverse: set comment to column: "original_cep" on table: "prices"
COMMENT ON COLUMN "prices" ."original_cep" IS '';
-- reverse: modify "prices" table
ALTER TABLE "prices" DROP COLUMN "original_cep";
