-- modify "bills" table
ALTER TABLE "bills" ADD COLUMN "target_symbol_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "symbol_id" on table: "bills"
COMMENT ON COLUMN "bills" ."symbol_id" IS '消费的外键币种 id';
-- set comment to column: "target_symbol_id" on table: "bills"
COMMENT ON COLUMN "bills" ."target_symbol_id" IS '获得的外键币种 id';
