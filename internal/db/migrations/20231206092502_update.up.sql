-- modify "users" table
ALTER TABLE "users" ADD COLUMN "nick_name" character varying NOT NULL DEFAULT '请设置昵称';
-- set comment to column: "nick_name" on table: "users"
COMMENT ON COLUMN "users" ."nick_name" IS '用户昵称';
