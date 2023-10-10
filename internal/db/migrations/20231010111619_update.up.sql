-- modify "users" table
ALTER TABLE "users" ADD COLUMN "pop_version" character varying NOT NULL DEFAULT '';
-- set comment to column: "pop_version" on table: "users"
COMMENT ON COLUMN "users" ."pop_version" IS '用户最新弹窗版本';
