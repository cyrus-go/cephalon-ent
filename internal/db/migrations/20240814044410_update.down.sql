-- reverse: drop index "usermodel_user_id_model_id" from table: "user_models"
CREATE UNIQUE INDEX "usermodel_user_id_model_id" ON "user_models" ("user_id", "model_id");
