-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "create_way" character varying NOT NULL DEFAULT 'unknown';
-- set comment to column: "create_way" on table: "missions"
COMMENT ON COLUMN "missions" ."create_way" IS '任务创建方式，user：用户自己通过官网页面创建，load_balance：用户通过创建 Loadbalance 自动创建';
