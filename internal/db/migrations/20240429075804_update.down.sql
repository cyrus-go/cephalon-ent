-- reverse: set comment to column: "weight" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."weight" IS '';
-- reverse: set comment to column: "type" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."type" IS '';
-- reverse: set comment to column: "category" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."category" IS '';
-- reverse: set comment to column: "deleted_at" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."id" IS '';
-- reverse: set comment to table: "mission_categories"
COMMENT ON TABLE "mission_categories" IS '';
-- reverse: create index "mission_categories_category_key" to table: "mission_categories"
DROP INDEX "mission_categories_category_key";
-- reverse: create "mission_categories" table
DROP TABLE "mission_categories";
