-- create "frpc_infos" table
CREATE TABLE "frpc_infos" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "tag" character varying NOT NULL DEFAULT '', "type" character varying NOT NULL DEFAULT '', "local_ip" character varying NOT NULL DEFAULT '', "local_port" bigint NOT NULL DEFAULT 0, "remote_port" bigint NOT NULL DEFAULT 0, "device_id" bigint NOT NULL DEFAULT 0, "frps_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "tag" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."tag" IS 'ini 文件客户端 tag';
-- set comment to column: "type" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."type" IS 'frpc 通讯类型';
-- set comment to column: "local_ip" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."local_ip" IS 'frpc 本地地址';
-- set comment to column: "local_port" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."local_port" IS 'frpc 本地要使用端口';
-- set comment to column: "remote_port" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."remote_port" IS 'frpc 本地要使用端口对应的远程端口';
-- set comment to column: "device_id" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."device_id" IS '外键设备 id';
-- set comment to column: "frps_id" on table: "frpc_infos"
COMMENT ON COLUMN "frpc_infos" ."frps_id" IS '外键 frps id';
-- create "frps_infos" table
CREATE TABLE "frps_infos" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "tag" character varying NOT NULL DEFAULT '', "server_addr" character varying NOT NULL DEFAULT '', "server_port" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "tag" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."tag" IS 'ini 文件服务端 tag';
-- set comment to column: "server_addr" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."server_addr" IS 'frps 服务地址';
-- set comment to column: "server_port" on table: "frps_infos"
COMMENT ON COLUMN "frps_infos" ."server_port" IS 'frps 服务端口';
