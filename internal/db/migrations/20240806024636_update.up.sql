-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "timed_shutdown" timestamptz NULL;
-- set comment to column: "timed_shutdown" on table: "missions"
COMMENT ON COLUMN "missions" ."timed_shutdown" IS '任务定时关机时间';
