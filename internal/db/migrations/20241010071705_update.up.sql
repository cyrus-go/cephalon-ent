-- modify "income_manages" table
ALTER TABLE "income_manages" ADD COLUMN "symbol_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "symbol_id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."symbol_id" IS '消费的外键币种 id';
