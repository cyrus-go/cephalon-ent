-- create "model_stars" table
CREATE TABLE "model_stars" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'unknown', "user_id" bigint NOT NULL, "model_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- create index "modelstar_user_id_model_id" to table: "model_stars"
CREATE UNIQUE INDEX "modelstar_user_id_model_id" ON "model_stars" ("user_id", "model_id");
-- set comment to column: "id" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "status" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."status" IS '收藏状态';
-- set comment to column: "user_id" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."user_id" IS '用户ID';
-- set comment to column: "model_id" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."model_id" IS '模型ID';
