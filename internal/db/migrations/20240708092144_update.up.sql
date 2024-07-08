-- modify "surveys" table
ALTER TABLE "surveys" ADD COLUMN "started_at" timestamptz NULL, ADD COLUMN "ended_at" timestamptz NULL, ADD COLUMN "sort_num" bigint NOT NULL DEFAULT 1, ADD COLUMN "group" character varying NOT NULL DEFAULT '';
-- set comment to column: "started_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."started_at" IS '填写问卷开始的时间';
-- set comment to column: "ended_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."ended_at" IS '填写问卷结束的时间';
-- set comment to column: "sort_num" on table: "surveys"
COMMENT ON COLUMN "surveys" ."sort_num" IS '分组排序序列号';
-- set comment to column: "group" on table: "surveys"
COMMENT ON COLUMN "surveys" ."group" IS '问卷分组（自定义，可以为空），同组问卷可以根据序号强关联';
