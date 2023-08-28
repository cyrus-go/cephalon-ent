-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "resp_status_code" integer NOT NULL DEFAULT 0, ADD COLUMN "resp_body" character varying NOT NULL DEFAULT '', ADD COLUMN "sd_api" character varying NOT NULL DEFAULT '';
-- set comment to column: "resp_status_code" on table: "missions"
COMMENT ON COLUMN "missions" ."resp_status_code" IS '内部功能返回码';
-- set comment to column: "resp_body" on table: "missions"
COMMENT ON COLUMN "missions" ."resp_body" IS '返回内容体 json 序列化为 string';
-- set comment to column: "sd_api" on table: "missions"
COMMENT ON COLUMN "missions" ."sd_api" IS '当 type 为 sd_api 时使用，为转发的 sd 接口功能';
