-- modify "bills" table
ALTER TABLE "bills" ADD COLUMN "profit_symbol_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "profit_symbol_id" on table: "bills"
COMMENT ON COLUMN "bills" ."profit_symbol_id" IS '外键分润币种 id';
-- modify "mission_orders" table
ALTER TABLE "mission_orders" ALTER COLUMN "last_settled_at" SET DEFAULT '0001-01-01 00:00:00.0000000 +00:00';
