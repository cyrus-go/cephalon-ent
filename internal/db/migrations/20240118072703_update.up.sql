-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "closed_at" timestamptz NULL;
-- set comment to column: "closed_at" on table: "missions"
COMMENT ON COLUMN "missions" ."closed_at" IS '用戶关闭任务时间';
