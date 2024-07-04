-- create "user_close_records" table
CREATE TABLE "user_close_records" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "registered_at" timestamptz NOT NULL, "closed_at" timestamptz NOT NULL, "type" character varying NOT NULL DEFAULT 'unknown', "user_id" bigint NOT NULL DEFAULT 0, "operate_user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "user_close_records"
COMMENT ON TABLE "user_close_records" IS '用户注销记录表';
-- set comment to column: "id" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "registered_at" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."registered_at" IS '本次注销时的注册时间';
-- set comment to column: "closed_at" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."closed_at" IS '本次注销的时间';
-- set comment to column: "type" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."type" IS '注销类型，用户自己注销或管理人员注销等';
-- set comment to column: "user_id" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."user_id" IS '用户 id';
-- set comment to column: "operate_user_id" on table: "user_close_records"
COMMENT ON COLUMN "user_close_records" ."operate_user_id" IS '操作人用户 id（只有管理人员注销时才有值）';
