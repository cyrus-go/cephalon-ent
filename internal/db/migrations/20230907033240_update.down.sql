-- reverse: set comment to column: "server_port" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."server_port" IS '';
-- reverse: set comment to column: "server_addr" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."server_addr" IS '';
-- reverse: set comment to column: "tag" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."tag" IS '';
-- reverse: create "frps_infos" table
DROP TABLE "frps_infos";
-- reverse: set comment to column: "frps_id" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."frps_id" IS '';
-- reverse: set comment to column: "device_id" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."device_id" IS '';
-- reverse: set comment to column: "remote_port" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."remote_port" IS '';
-- reverse: set comment to column: "local_port" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."local_port" IS '';
-- reverse: set comment to column: "local_ip" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."local_ip" IS '';
-- reverse: set comment to column: "type" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."type" IS '';
-- reverse: set comment to column: "tag" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."tag" IS '';
-- reverse: create "frpc_infos" table
DROP TABLE "frpc_infos";
