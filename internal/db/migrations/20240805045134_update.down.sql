-- reverse: set comment to column: "model_id" on table: "user_models"
COMMENT ON COLUMN "user_models" ."model_id" IS '';
-- reverse: set comment to column: "user_id" on table: "user_models"
COMMENT ON COLUMN "user_models" ."user_id" IS '';
-- reverse: set comment to column: "status" on table: "user_models"
COMMENT ON COLUMN "user_models" ."status" IS '';
-- reverse: set comment to column: "relation" on table: "user_models"
COMMENT ON COLUMN "user_models" ."relation" IS '';
-- reverse: set comment to column: "deleted_at" on table: "user_models"
COMMENT ON COLUMN "user_models" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "user_models"
COMMENT ON COLUMN "user_models" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "user_models"
COMMENT ON COLUMN "user_models" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "user_models"
COMMENT ON COLUMN "user_models" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "user_models"
COMMENT ON COLUMN "user_models" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "user_models"
COMMENT ON COLUMN "user_models" ."id" IS '';
-- reverse: create index "usermodel_user_id_model_id" to table: "user_models"
DROP INDEX "usermodel_user_id_model_id";
-- reverse: create "user_models" table
DROP TABLE "user_models";
