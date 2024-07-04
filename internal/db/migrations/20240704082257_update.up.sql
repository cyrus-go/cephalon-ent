-- modify "users" table
ALTER TABLE "users" ADD COLUMN "re_register_at" timestamptz NULL;
-- set comment to column: "re_register_at" on table: "users"
COMMENT ON COLUMN "users" ."re_register_at" IS '用户注销之后重新注册的时间';
