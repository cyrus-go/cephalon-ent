-- modify "frps_infos" table
ALTER TABLE "frps_infos" ADD COLUMN "type" character varying NOT NULL DEFAULT '';
-- set comment to column: "type" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."type" IS '类型：public、private';
