-- modify "users" table
ALTER TABLE "users" ADD COLUMN "github_id" bigint NOT NULL DEFAULT 0;
-- create index "user_github_id" to table: "users"
CREATE INDEX "user_github_id" ON "users" ("github_id");
-- set comment to column: "github_id" on table: "users"
COMMENT ON COLUMN "users" ."github_id" IS '第三方登录github id';
