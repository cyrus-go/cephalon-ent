-- modify "users" table
ALTER TABLE "users" ADD COLUMN "google_id" character varying NOT NULL DEFAULT '';
-- set comment to column: "google_id" on table: "users"
COMMENT ON COLUMN "users" ."google_id" IS '第三方登录google id';
