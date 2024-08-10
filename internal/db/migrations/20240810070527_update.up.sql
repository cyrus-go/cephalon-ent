-- modify "model_prices" table
ALTER TABLE "model_prices" ADD COLUMN "token_per_cep" bigint NOT NULL DEFAULT 0;
-- set comment to column: "token_per_cep" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."token_per_cep" IS '上面cep价格对应的token数量';
