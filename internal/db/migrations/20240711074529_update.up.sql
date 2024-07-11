-- modify "surveys" table
ALTER TABLE "surveys" ADD COLUMN "background_image" character varying NOT NULL DEFAULT '';
-- set comment to column: "background_image" on table: "surveys"
COMMENT ON COLUMN "surveys" ."background_image" IS '问卷背景图';
