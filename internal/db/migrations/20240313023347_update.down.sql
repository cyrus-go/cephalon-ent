-- reverse: set comment to column: "baidu_access_token" on table: "users"
COMMENT ON COLUMN "users" ."baidu_access_token" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "baidu_access_token";
