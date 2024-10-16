-- reverse: set comment to column: "status" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."status" IS '';
-- reverse: set comment to column: "version" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."version" IS '';
-- reverse: set comment to column: "config_url" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."config_url" IS '';
-- reverse: set comment to column: "client_url" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."client_url" IS '';
-- reverse: set comment to column: "deleted_at" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "client_versions"
COMMENT ON COLUMN "client_versions" ."id" IS '';
-- reverse: set comment to table: "client_versions"
COMMENT ON TABLE "client_versions" IS '';
-- reverse: create index "clientversion_status" to table: "client_versions"
DROP INDEX "clientversion_status";
-- reverse: create index "client_versions_version_key" to table: "client_versions"
DROP INDEX "client_versions_version_key";
-- reverse: create "client_versions" table
DROP TABLE "client_versions";
-- reverse: set comment to column: "mission_tag" on table: "users"
COMMENT ON COLUMN "users" ."mission_tag" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "mission_tag";
