-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "version" character varying NOT NULL DEFAULT '无版本号';
-- set comment to column: "version" on table: "devices"
COMMENT ON COLUMN "devices" ."version" IS '设备版本';
