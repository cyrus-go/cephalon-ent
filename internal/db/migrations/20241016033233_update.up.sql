-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "mission_tag" character varying NOT NULL DEFAULT 'no';
-- set comment to column: "mission_tag" on table: "devices"
COMMENT ON COLUMN "devices" ."mission_tag" IS '可接特殊任务类型标签';
