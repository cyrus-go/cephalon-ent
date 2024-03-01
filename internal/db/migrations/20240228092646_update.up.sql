-- create "cloud_files" table
CREATE TABLE "cloud_files" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', "icon" character varying NOT NULL DEFAULT '', "size" bigint NOT NULL DEFAULT 0, "md5" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "cloudfile_user_id" to table: "cloud_files"
CREATE INDEX "cloudfile_user_id" ON "cloud_files" ("user_id");
-- set comment to column: "id" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "name" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."name" IS '文件名';
-- set comment to column: "icon" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."icon" IS '文件图标';
-- set comment to column: "size" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."size" IS '文件大小';
-- set comment to column: "md5" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."md5" IS 'md5';
-- set comment to column: "user_id" on table: "cloud_files"
COMMENT ON COLUMN "cloud_files" ."user_id" IS '外键用户id';
