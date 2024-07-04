-- reverse: set comment to column: "re_register_at" on table: "users"
COMMENT ON COLUMN "users" ."re_register_at" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "re_register_at";
