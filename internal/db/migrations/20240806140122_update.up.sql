-- create "invoke_model_orders" table
CREATE TABLE "invoke_model_orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "invoke_type" character varying NOT NULL DEFAULT 'unknown', "invoke_times" bigint NOT NULL DEFAULT 0, "input_token_cost" bigint NOT NULL DEFAULT 0, "output_token_cost" bigint NOT NULL DEFAULT 0, "input_cep_cost" bigint NOT NULL DEFAULT 0, "output_cep_cost" bigint NOT NULL DEFAULT 0, "api_token_id" bigint NOT NULL, "model_id" bigint NOT NULL, "user_id" bigint NOT NULL, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "invoke_type" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."invoke_type" IS '调用类型';
-- set comment to column: "invoke_times" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."invoke_times" IS '调用次数';
-- set comment to column: "input_token_cost" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."input_token_cost" IS '输入token消耗';
-- set comment to column: "output_token_cost" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."output_token_cost" IS '输出token消耗';
-- set comment to column: "input_cep_cost" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."input_cep_cost" IS '输入cep消耗';
-- set comment to column: "output_cep_cost" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."output_cep_cost" IS '输出cep消耗';
-- set comment to column: "api_token_id" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."api_token_id" IS 'API Token ID';
-- set comment to column: "model_id" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."model_id" IS '模型ID';
-- set comment to column: "user_id" on table: "invoke_model_orders"
COMMENT ON COLUMN "invoke_model_orders" ."user_id" IS '用户ID';
