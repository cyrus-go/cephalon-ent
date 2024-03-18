-- modify "frps_infos" table
ALTER TABLE "frps_infos" ADD COLUMN "domain" character varying NOT NULL DEFAULT '';
-- set comment to column: "domain" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."domain" IS '域名';
