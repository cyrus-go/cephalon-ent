-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "remark" character varying NOT NULL DEFAULT '';
-- set comment to column: "remark" on table: "missions"
COMMENT ON COLUMN "missions" ."remark" IS '备注信息';
