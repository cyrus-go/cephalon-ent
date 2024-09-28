-- drop index "user_github_id" from table: "users"
DROP INDEX "user_github_id";
-- create index "user_phone_deleted_at_email_github_id_google_id" to table: "users"
CREATE UNIQUE INDEX "user_phone_deleted_at_email_github_id_google_id" ON "users" ("phone", "deleted_at", "email", "github_id", "google_id");
