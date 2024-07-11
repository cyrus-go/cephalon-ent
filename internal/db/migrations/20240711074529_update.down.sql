-- reverse: set comment to column: "background_image" on table: "surveys"
COMMENT ON COLUMN "surveys" ."background_image" IS '';
-- reverse: modify "surveys" table
ALTER TABLE "surveys" DROP COLUMN "background_image";
