-- modify "frpc_infos" table
ALTER TABLE "frpc_infos" ADD COLUMN "is_using" boolean NOT NULL DEFAULT false;
-- set comment to column: "is_using" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."is_using" IS '端口是否已经在使用';
