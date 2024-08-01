-- create "api_tokens" table
CREATE TABLE "api_tokens" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', "token" character varying NOT NULL DEFAULT '', "status" character varying NOT NULL DEFAULT 'unknown', "user_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "name" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."name" IS 'token 名称';
-- set comment to column: "token" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."token" IS 'token 内容';
-- set comment to column: "status" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."status" IS 'token 状态';
-- set comment to column: "user_id" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."user_id" IS '用户ID';
-- create "models" table
CREATE TABLE "models" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', "author" character varying NOT NULL DEFAULT '', "model_type" character varying NOT NULL DEFAULT 'unknown', "model_status" character varying NOT NULL DEFAULT 'unknown', "is_official" boolean NOT NULL DEFAULT false, "total_usage" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "models"
COMMENT ON COLUMN "models" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "models"
COMMENT ON COLUMN "models" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "models"
COMMENT ON COLUMN "models" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "models"
COMMENT ON COLUMN "models" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "models"
COMMENT ON COLUMN "models" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "models"
COMMENT ON COLUMN "models" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "name" on table: "models"
COMMENT ON COLUMN "models" ."name" IS '模型名称';
-- set comment to column: "author" on table: "models"
COMMENT ON COLUMN "models" ."author" IS '模型作者';
-- set comment to column: "model_type" on table: "models"
COMMENT ON COLUMN "models" ."model_type" IS '模型类型';
-- set comment to column: "model_status" on table: "models"
COMMENT ON COLUMN "models" ."model_status" IS '模型状态';
-- set comment to column: "is_official" on table: "models"
COMMENT ON COLUMN "models" ."is_official" IS '是否为官方模型';
-- set comment to column: "total_usage" on table: "models"
COMMENT ON COLUMN "models" ."total_usage" IS '模型的总使用次数';
-- create "model_prices" table
CREATE TABLE "model_prices" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "invoke_type" character varying NOT NULL DEFAULT 'unknown', "gpu_version" character varying NOT NULL DEFAULT 'unknown', "input_gpu_price" bigint NOT NULL DEFAULT 0, "output_gpu_price" bigint NOT NULL DEFAULT 0, "input_model_price" bigint NOT NULL DEFAULT 0, "output_model_price" bigint NOT NULL DEFAULT 0, "model_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "invoke_type" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."invoke_type" IS '调用类型';
-- set comment to column: "gpu_version" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."gpu_version" IS 'GPU 版本';
-- set comment to column: "input_gpu_price" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."input_gpu_price" IS '输入算力价格';
-- set comment to column: "output_gpu_price" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."output_gpu_price" IS '输出算力价格';
-- set comment to column: "input_model_price" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."input_model_price" IS '输入模型使用价格';
-- set comment to column: "output_model_price" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."output_model_price" IS '输出模型使用价格';
-- set comment to column: "model_id" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."model_id" IS '模型ID';
-- create "modle_stars" table
CREATE TABLE "modle_stars" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'unknown', "user_id" bigint NOT NULL, "model_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- create index "modlestar_user_id_model_id" to table: "modle_stars"
CREATE UNIQUE INDEX "modlestar_user_id_model_id" ON "modle_stars" ("user_id", "model_id");
-- set comment to column: "id" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "status" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."status" IS '收藏状态';
-- set comment to column: "user_id" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."user_id" IS '用户ID';
-- set comment to column: "model_id" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."model_id" IS '模型ID';
