-- modify "models" table
ALTER TABLE "models" ALTER COLUMN "model_status" SET DEFAULT 'init', ADD COLUMN "description" character varying NOT NULL DEFAULT '';
-- set comment to column: "description" on table: "models"
COMMENT ON COLUMN "models" ."description" IS '模型描述';
