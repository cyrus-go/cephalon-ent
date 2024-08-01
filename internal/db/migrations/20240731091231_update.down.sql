-- reverse: set comment to column: "model_id" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."model_id" IS '';
-- reverse: set comment to column: "user_id" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."user_id" IS '';
-- reverse: set comment to column: "status" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."status" IS '';
-- reverse: set comment to column: "deleted_at" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "modle_stars"
COMMENT ON COLUMN "modle_stars" ."id" IS '';
-- reverse: create index "modlestar_user_id_model_id" to table: "modle_stars"
DROP INDEX "modlestar_user_id_model_id";
-- reverse: create "modle_stars" table
DROP TABLE "modle_stars";
-- reverse: set comment to column: "model_id" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."model_id" IS '';
-- reverse: set comment to column: "output_model_price" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."output_model_price" IS '';
-- reverse: set comment to column: "input_model_price" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."input_model_price" IS '';
-- reverse: set comment to column: "output_gpu_price" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."output_gpu_price" IS '';
-- reverse: set comment to column: "input_gpu_price" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."input_gpu_price" IS '';
-- reverse: set comment to column: "gpu_version" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."gpu_version" IS '';
-- reverse: set comment to column: "invoke_type" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."invoke_type" IS '';
-- reverse: set comment to column: "deleted_at" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "model_prices"
COMMENT ON COLUMN "model_prices" ."id" IS '';
-- reverse: create "model_prices" table
DROP TABLE "model_prices";
-- reverse: set comment to column: "total_usage" on table: "models"
COMMENT ON COLUMN "models" ."total_usage" IS '';
-- reverse: set comment to column: "is_official" on table: "models"
COMMENT ON COLUMN "models" ."is_official" IS '';
-- reverse: set comment to column: "model_status" on table: "models"
COMMENT ON COLUMN "models" ."model_status" IS '';
-- reverse: set comment to column: "model_type" on table: "models"
COMMENT ON COLUMN "models" ."model_type" IS '';
-- reverse: set comment to column: "author" on table: "models"
COMMENT ON COLUMN "models" ."author" IS '';
-- reverse: set comment to column: "name" on table: "models"
COMMENT ON COLUMN "models" ."name" IS '';
-- reverse: set comment to column: "deleted_at" on table: "models"
COMMENT ON COLUMN "models" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "models"
COMMENT ON COLUMN "models" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "models"
COMMENT ON COLUMN "models" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "models"
COMMENT ON COLUMN "models" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "models"
COMMENT ON COLUMN "models" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "models"
COMMENT ON COLUMN "models" ."id" IS '';
-- reverse: create "models" table
DROP TABLE "models";
-- reverse: set comment to column: "user_id" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."user_id" IS '';
-- reverse: set comment to column: "status" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."status" IS '';
-- reverse: set comment to column: "token" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."token" IS '';
-- reverse: set comment to column: "name" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."name" IS '';
-- reverse: set comment to column: "deleted_at" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "api_tokens"
COMMENT ON COLUMN "api_tokens" ."id" IS '';
-- reverse: create "api_tokens" table
DROP TABLE "api_tokens";
