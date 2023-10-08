-- reverse: set comment to column: "symbol_id" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."symbol_id" IS '';
-- reverse: modify "renewal_agreements" table
ALTER TABLE "renewal_agreements" DROP COLUMN "symbol_id";
