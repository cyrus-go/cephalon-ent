-- create "user_models" table
CREATE TABLE "user_models" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "relation" character varying NOT NULL DEFAULT 'unknown', "status" character varying NOT NULL DEFAULT 'unknown', "user_id" bigint NOT NULL, "model_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- create index "usermodel_user_id_model_id" to table: "user_models"
CREATE UNIQUE INDEX "usermodel_user_id_model_id" ON "user_models" ("user_id", "model_id");
-- set comment to column: "id" on table: "user_models"
COMMENT ON COLUMN "user_models" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "user_models"
COMMENT ON COLUMN "user_models" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "user_models"
COMMENT ON COLUMN "user_models" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "user_models"
COMMENT ON COLUMN "user_models" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "user_models"
COMMENT ON COLUMN "user_models" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "user_models"
COMMENT ON COLUMN "user_models" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "relation" on table: "user_models"
COMMENT ON COLUMN "user_models" ."relation" IS '关系';
-- set comment to column: "status" on table: "user_models"
COMMENT ON COLUMN "user_models" ."status" IS '状态';
-- set comment to column: "user_id" on table: "user_models"
COMMENT ON COLUMN "user_models" ."user_id" IS '用户ID';
-- set comment to column: "model_id" on table: "user_models"
COMMENT ON COLUMN "user_models" ."model_id" IS '模型ID';
