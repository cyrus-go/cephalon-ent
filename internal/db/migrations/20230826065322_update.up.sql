-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "unit_cep" bigint NOT NULL DEFAULT 0;
-- set comment to column: "unit_cep" on table: "missions"
COMMENT ON COLUMN "missions" ."unit_cep" IS '任务单价，按次就是 unit_cep/次，按时就是 unit_cep/分钟';
-- modify "prices" table
ALTER TABLE "prices" ADD COLUMN "started_at" timestamptz NULL, ADD COLUMN "finished_at" timestamptz NULL;
-- set comment to column: "started_at" on table: "prices"
COMMENT ON COLUMN "prices" ."started_at" IS '价格有效时间开始，为空表示永久有效';
-- set comment to column: "finished_at" on table: "prices"
COMMENT ON COLUMN "prices" ."finished_at" IS '价格有效时间结束，为空表示永久有效';
