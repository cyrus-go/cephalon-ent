-- modify "users" table
ALTER TABLE "users" ADD COLUMN "baidu_refresh_token" character varying NOT NULL DEFAULT '';
-- set comment to column: "baidu_refresh_token" on table: "users"
COMMENT ON COLUMN "users" ."baidu_refresh_token" IS '百度网盘刷新 token';
