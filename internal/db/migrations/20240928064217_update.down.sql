-- reverse: create index "user_phone_deleted_at_email_github_id_google_id" to table: "users"
DROP INDEX "user_phone_deleted_at_email_github_id_google_id";
-- reverse: drop index "user_github_id" from table: "users"
CREATE INDEX "user_github_id" ON "users" ("github_id");
