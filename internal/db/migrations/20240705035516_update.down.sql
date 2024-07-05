-- reverse: set comment to column: "user_id" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."user_id" IS '';
-- reverse: set comment to column: "survey_id" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."survey_id" IS '';
-- reverse: set comment to column: "deleted_at" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."id" IS '';
-- reverse: set comment to table: "survey_responses"
COMMENT ON TABLE "survey_responses" IS '';
-- reverse: create "survey_responses" table
DROP TABLE "survey_responses";
-- reverse: set comment to column: "survey_id" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."survey_id" IS '';
-- reverse: set comment to column: "options" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."options" IS '';
-- reverse: set comment to column: "type" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."type" IS '';
-- reverse: set comment to column: "text" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."text" IS '';
-- reverse: set comment to column: "deleted_at" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."id" IS '';
-- reverse: set comment to table: "survey_questions"
COMMENT ON TABLE "survey_questions" IS '';
-- reverse: create "survey_questions" table
DROP TABLE "survey_questions";
-- reverse: set comment to column: "survey_response_id" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."survey_response_id" IS '';
-- reverse: set comment to column: "survey_question_id" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."survey_question_id" IS '';
-- reverse: set comment to column: "survey_answer" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."survey_answer" IS '';
-- reverse: set comment to column: "deleted_at" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."id" IS '';
-- reverse: set comment to table: "survey_answers"
COMMENT ON TABLE "survey_answers" IS '';
-- reverse: create "survey_answers" table
DROP TABLE "survey_answers";
-- reverse: set comment to column: "title" on table: "surveys"
COMMENT ON COLUMN "surveys" ."title" IS '';
-- reverse: set comment to column: "deleted_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "surveys"
COMMENT ON COLUMN "surveys" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "surveys"
COMMENT ON COLUMN "surveys" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "surveys"
COMMENT ON COLUMN "surveys" ."id" IS '';
-- reverse: set comment to table: "surveys"
COMMENT ON TABLE "surveys" IS '';
-- reverse: create "surveys" table
DROP TABLE "surveys";
