-- reverse: set comment to column: "finished_at" on table: "missions"
COMMENT ON COLUMN "missions" ."finished_at" IS '';
-- reverse: set comment to column: "expired_at" on table: "missions"
COMMENT ON COLUMN "missions" ."expired_at" IS '任务到期时间';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "finished_at";
