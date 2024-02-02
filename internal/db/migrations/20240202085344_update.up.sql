-- modify "prizes" table
ALTER TABLE "prizes" ADD COLUMN "type" character varying NOT NULL DEFAULT 'unknow', ADD COLUMN "cep_amount" bigint NOT NULL DEFAULT 0;
-- set comment to column: "type" on table: "prizes"
COMMENT ON COLUMN "prizes" ."type" IS '类型';
-- set comment to column: "cep_amount" on table: "prizes"
COMMENT ON COLUMN "prizes" ."cep_amount" IS '类型为 get_cep 时，cep 的数量';
