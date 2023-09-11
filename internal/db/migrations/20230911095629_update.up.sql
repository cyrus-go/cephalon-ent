-- modify "frps_infos" table
ALTER TABLE "frps_infos" ADD COLUMN "authentication_method" character varying NOT NULL DEFAULT '', ADD COLUMN "token" character varying NOT NULL DEFAULT '';
-- set comment to column: "authentication_method" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."authentication_method" IS 'frps 认证方式';
-- set comment to column: "token" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."token" IS 'frps 认证 token';
