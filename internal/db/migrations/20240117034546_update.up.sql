-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "close_way" character varying NOT NULL DEFAULT 'unknown';
-- set comment to column: "close_way" on table: "missions"
COMMENT ON COLUMN "missions" ."close_way" IS '任务关闭方式，user：用户自己关闭，balance_not_enough：余额不足自动关闭';
