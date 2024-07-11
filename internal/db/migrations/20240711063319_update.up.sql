-- modify "surveys" table
ALTER TABLE "surveys" ADD COLUMN "hint" character varying NOT NULL DEFAULT '';
-- set comment to column: "hint" on table: "surveys"
COMMENT ON COLUMN "surveys" ."hint" IS '问卷提示信息';
