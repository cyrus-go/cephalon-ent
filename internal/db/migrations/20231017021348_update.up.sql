-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "call_back_data" bytea NULL, ADD COLUMN "username" character varying NOT NULL DEFAULT 'admin', ADD COLUMN "password" character varying NOT NULL DEFAULT 'cephalon';
-- set comment to column: "call_back_data" on table: "missions"
COMMENT ON COLUMN "missions" ."call_back_data" IS '回调时带上的参数，需 json 反序列化后才可使用，所以没有直接 json 序列化字段 (sensitive)';
-- set comment to column: "username" on table: "missions"
COMMENT ON COLUMN "missions" ."username" IS '某些任务会使用到的验证用户名';
-- set comment to column: "password" on table: "missions"
COMMENT ON COLUMN "missions" ."password" IS '某些任务会使用到的验证密码';
