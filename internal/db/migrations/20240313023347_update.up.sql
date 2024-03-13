-- modify "users" table
ALTER TABLE "users" ADD COLUMN "baidu_access_token" character varying NOT NULL DEFAULT '';
-- set comment to column: "baidu_access_token" on table: "users"
COMMENT ON COLUMN "users" ."baidu_access_token" IS '百度网盘 token';
