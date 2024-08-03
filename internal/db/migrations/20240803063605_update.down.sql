-- reverse: set comment to column: "description" on table: "models"
COMMENT ON COLUMN "models" ."description" IS '';
-- reverse: modify "models" table
ALTER TABLE "models" DROP COLUMN "description", ALTER COLUMN "model_status" SET DEFAULT 'unknown';
