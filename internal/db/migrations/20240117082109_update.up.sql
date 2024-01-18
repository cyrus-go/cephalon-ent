-- modify "prices" table
ALTER TABLE "prices" ADD COLUMN "original_cep" bigint NOT NULL DEFAULT 0;
-- set comment to column: "original_cep" on table: "prices"
COMMENT ON COLUMN "prices" ."original_cep" IS '任务原价';
