-- modify "mission_orders" table
ALTER TABLE "mission_orders" ADD COLUMN "plan_started_at" timestamptz NULL, ADD COLUMN "plan_finished_at" timestamptz NULL;
-- set comment to column: "plan_started_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."plan_started_at" IS '任务计划开始时间（包时）';
-- set comment to column: "plan_finished_at" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."plan_finished_at" IS '任务计划结束时间（包时）';
-- create "login_records" table
CREATE TABLE "login_records" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "ua" character varying NOT NULL DEFAULT '', "ip" character varying NOT NULL DEFAULT '', "way" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "login_records"
COMMENT ON TABLE "login_records" IS '登录记录，记录用户登录的一些信息';
-- set comment to column: "ua" on table: "login_records"
COMMENT ON COLUMN "login_records" ."ua" IS '用户登录时的 user-agent';
-- set comment to column: "ip" on table: "login_records"
COMMENT ON COLUMN "login_records" ."ip" IS '用户登录时的 ip 地址';
-- set comment to column: "way" on table: "login_records"
COMMENT ON COLUMN "login_records" ."way" IS '用户登录时的登录方式，比如手机号验证码等';
-- set comment to column: "user_id" on table: "login_records"
COMMENT ON COLUMN "login_records" ."user_id" IS '用户 id，外键关联用户';
