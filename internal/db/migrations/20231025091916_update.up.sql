-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "finished_at" timestamptz NULL;
-- set comment to column: "expired_at" on table: "missions"
COMMENT ON COLUMN "missions" ."expired_at" IS '任务到期时间（包时任务才有）';
-- set comment to column: "finished_at" on table: "missions"
COMMENT ON COLUMN "missions" ."finished_at" IS '任务结束时间';
