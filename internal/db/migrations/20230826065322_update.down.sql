-- reverse: set comment to column: "finished_at" on table: "prices"
COMMENT ON COLUMN "prices" ."finished_at" IS '';
-- reverse: set comment to column: "started_at" on table: "prices"
COMMENT ON COLUMN "prices" ."started_at" IS '';
-- reverse: modify "prices" table
ALTER TABLE "prices" DROP COLUMN "finished_at", DROP COLUMN "started_at";
-- reverse: set comment to column: "unit_cep" on table: "missions"
COMMENT ON COLUMN "missions" ."unit_cep" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "unit_cep";
