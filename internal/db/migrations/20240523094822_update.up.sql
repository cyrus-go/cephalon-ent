-- modify "users" table
ALTER TABLE "users" ADD COLUMN "bound_at" timestamptz NULL;
-- set comment to column: "bound_at" on table: "users"
COMMENT ON COLUMN "users" ."bound_at" IS '用户绑定邀请码的时间';
