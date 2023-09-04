-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "second_hmac_key" character varying NOT NULL DEFAULT '';
-- set comment to column: "second_hmac_key" on table: "missions"
COMMENT ON COLUMN "missions" ."second_hmac_key" IS '创建任务时使用了的 二级 hmac_key';
