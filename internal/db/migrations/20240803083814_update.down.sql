-- reverse: set comment to column: "model_id" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."model_id" IS '';
-- reverse: set comment to column: "user_id" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."user_id" IS '';
-- reverse: set comment to column: "status" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."status" IS '';
-- reverse: set comment to column: "deleted_at" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "model_stars"
COMMENT ON COLUMN "model_stars" ."id" IS '';
-- reverse: create index "modelstar_user_id_model_id" to table: "model_stars"
DROP INDEX "modelstar_user_id_model_id";
-- reverse: create "model_stars" table
DROP TABLE "model_stars";
