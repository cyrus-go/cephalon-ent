-- modify "users" table
ALTER TABLE "users" ADD COLUMN "mission_tag" character varying NOT NULL DEFAULT 'no';
-- set comment to column: "mission_tag" on table: "users"
COMMENT ON COLUMN "users" ."mission_tag" IS '可跳过验证启动特殊任务类型标签';
-- create "client_versions" table
CREATE TABLE "client_versions" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "client_url" character varying NOT NULL DEFAULT '', "config_url" character varying NOT NULL DEFAULT '', "version" character varying NOT NULL, "status" character varying NOT NULL DEFAULT 'disable', PRIMARY KEY ("id"));
-- create index "client_versions_version_key" to table: "client_versions"
CREATE UNIQUE INDEX "client_versions_version_key" ON "client_versions" ("version");
-- create index "clientversion_status" to table: "client_versions"
CREATE INDEX "clientversion_status" ON "client_versions" ("status");
-- set comment to table: "client_versions"
COMMENT ON TABLE "client_versions" IS '客户端版本，OTA 服务端功能支持';
-- set comment to column: "id" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "client_url" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."client_url" IS '客户端文件地址';
-- set comment to column: "config_url" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."config_url" IS '主配置文件地址';
-- set comment to column: "version" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."version" IS '版本号';
-- set comment to column: "status" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."status" IS '状态：只允许有一条数据的状态为启用（可被自动更新的版本）';
