-- modify "survey_responses" table
ALTER TABLE "survey_responses" ADD COLUMN "status" character varying NOT NULL DEFAULT 'pending', ADD COLUMN "approved_by" bigint NOT NULL DEFAULT 0;
-- set comment to column: "status" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."status" IS '调查问卷结果状态';
-- set comment to column: "approved_by" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."approved_by" IS '审批用户 ID';
-- modify "surveys" table
ALTER TABLE "surveys" ADD COLUMN "gift_cep_amount" bigint NOT NULL DEFAULT 0, ADD COLUMN "gift_type" character varying NOT NULL DEFAULT 'unknown', ADD COLUMN "desc" character varying NOT NULL DEFAULT '';
-- set comment to column: "gift_cep_amount" on table: "surveys"
COMMENT ON COLUMN "surveys" ."gift_cep_amount" IS '提交问卷赠送的脑力值数量';
-- set comment to column: "gift_type" on table: "surveys"
COMMENT ON COLUMN "surveys" ."gift_type" IS '问卷赠送类型，提交赠送或审批赠送等';
-- set comment to column: "desc" on table: "surveys"
COMMENT ON COLUMN "surveys" ."desc" IS '问卷描述信息';
