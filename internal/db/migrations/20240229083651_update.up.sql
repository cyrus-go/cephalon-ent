-- modify "users" table
ALTER TABLE "users" ADD COLUMN "cloud_space" bigint NOT NULL DEFAULT 0;
-- set comment to column: "cloud_space" on table: "users"
COMMENT ON COLUMN "users" ."cloud_space" IS '云盘空间';
