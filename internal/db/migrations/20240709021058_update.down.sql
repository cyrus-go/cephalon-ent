-- reverse: set comment to column: "desc" on table: "surveys"
COMMENT ON COLUMN "surveys" ."desc" IS '';
-- reverse: set comment to column: "gift_type" on table: "surveys"
COMMENT ON COLUMN "surveys" ."gift_type" IS '';
-- reverse: set comment to column: "gift_cep_amount" on table: "surveys"
COMMENT ON COLUMN "surveys" ."gift_cep_amount" IS '';
-- reverse: modify "surveys" table
ALTER TABLE "surveys" DROP COLUMN "desc", DROP COLUMN "gift_type", DROP COLUMN "gift_cep_amount";
-- reverse: set comment to column: "approved_by" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."approved_by" IS '';
-- reverse: set comment to column: "status" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."status" IS '';
-- reverse: modify "survey_responses" table
ALTER TABLE "survey_responses" DROP COLUMN "approved_by", DROP COLUMN "status";
