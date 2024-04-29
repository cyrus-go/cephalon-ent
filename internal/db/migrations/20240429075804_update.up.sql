-- create "mission_categories" table
CREATE TABLE "mission_categories" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "category" character varying NOT NULL, "type" character varying NOT NULL DEFAULT 'unknown', "weight" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "mission_categories_category_key" to table: "mission_categories"
CREATE UNIQUE INDEX "mission_categories_category_key" ON "mission_categories" ("category");
-- set comment to table: "mission_categories"
COMMENT ON TABLE "mission_categories" IS '任务大类，任务类型的最高抽象层，记录了所有任务大类';
-- set comment to column: "id" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "category" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."category" IS '任务大类';
-- set comment to column: "type" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."type" IS '任务大类类型，比如热门类型等';
-- set comment to column: "weight" on table: "mission_categories"
COMMENT ON COLUMN "mission_categories" ."weight" IS '权重（目前排序可以用到）';
