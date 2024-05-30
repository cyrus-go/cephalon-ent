-- reverse: set comment to column: "last_edited_user_id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."last_edited_user_id" IS '';
-- reverse: modify "income_manages" table
ALTER TABLE "income_manages" DROP COLUMN "last_edited_user_id";
