-- reverse: set comment to column: "baidu_refresh_token" on table: "users"
COMMENT ON COLUMN "users" ."baidu_refresh_token" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "baidu_refresh_token";
