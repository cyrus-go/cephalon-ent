-- reverse: set comment to column: "user_id" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."user_id" IS '';
-- reverse: set comment to column: "md5" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."md5" IS '';
-- reverse: set comment to column: "size" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."size" IS '';
-- reverse: set comment to column: "icon" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."icon" IS '';
-- reverse: set comment to column: "name" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."name" IS '';
-- reverse: set comment to column: "deleted_at" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."id" IS '';
-- reverse: create index "cloudfile_user_id" to table: "cloud_files"
DROP INDEX "cloudfile_user_id";
-- reverse: create "cloud_files" table
DROP TABLE "cloud_files";
