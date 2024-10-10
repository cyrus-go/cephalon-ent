-- modify "mission_load_balances" table
ALTER TABLE "mission_load_balances" ADD COLUMN "close_way" character varying NOT NULL DEFAULT 'unknown';
-- set comment to column: "close_way" on table: "mission_load_balances"
COMMENT ON COLUMN "mission_load_balances" ."close_way" IS '任务关闭方式，user：用户自己关闭，balance_not_enough：余额不足自动关闭';
