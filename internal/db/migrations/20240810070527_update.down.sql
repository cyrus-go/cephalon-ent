-- reverse: set comment to column: "token_per_cep" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."token_per_cep" IS '';
-- reverse: modify "model_prices" table
ALTER TABLE "model_prices" DROP COLUMN "token_per_cep";
