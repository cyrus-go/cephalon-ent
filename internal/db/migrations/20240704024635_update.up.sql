-- modify "users" table
ALTER TABLE "users" ADD COLUMN "user_status" character varying NOT NULL DEFAULT 'normal';
-- set comment to column: "user_status" on table: "users"
COMMENT ON COLUMN "users" ."user_status" IS '用户状态';
