-- modify "renewal_agreements" table
ALTER TABLE "renewal_agreements" ADD COLUMN "symbol_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "symbol_id" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."symbol_id" IS '币种 id';
