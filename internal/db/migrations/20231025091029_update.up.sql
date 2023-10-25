-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "started_at" timestamptz NULL, ADD COLUMN "expired_at" timestamptz NULL;
-- set comment to column: "started_at" on table: "missions"
COMMENT ON COLUMN "missions" ."started_at" IS '任务开始时间';
-- set comment to column: "expired_at" on table: "missions"
COMMENT ON COLUMN "missions" ."expired_at" IS '任务到期时间';
