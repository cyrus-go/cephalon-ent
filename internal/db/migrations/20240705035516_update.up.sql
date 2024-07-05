-- create "surveys" table
CREATE TABLE "surveys" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "title" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- set comment to table: "surveys"
COMMENT ON TABLE "surveys" IS '问卷表';
-- set comment to column: "id" on table: "surveys"
COMMENT ON COLUMN "surveys" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "surveys"
COMMENT ON COLUMN "surveys" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "surveys"
COMMENT ON COLUMN "surveys" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "surveys"
COMMENT ON COLUMN "surveys" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "title" on table: "surveys"
COMMENT ON COLUMN "surveys" ."title" IS '标题';
-- create "survey_answers" table
CREATE TABLE "survey_answers" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "survey_answer" text NOT NULL DEFAULT '', "survey_question_id" bigint NOT NULL DEFAULT 0, "survey_response_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "survey_answers"
COMMENT ON TABLE "survey_answers" IS '问卷问题答案表';
-- set comment to column: "id" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "survey_answer" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."survey_answer" IS '答案的内容';
-- set comment to column: "survey_question_id" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."survey_question_id" IS '问题 id';
-- set comment to column: "survey_response_id" on table: "survey_answers"
COMMENT ON COLUMN "survey_answers" ."survey_response_id" IS '问卷用户调查结果 ID，一个用户同一个问卷只能回答一次，这个问卷会有多个问题和答案';
-- create "survey_questions" table
CREATE TABLE "survey_questions" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "text" text NOT NULL DEFAULT '', "type" character varying NOT NULL DEFAULT 'unknown', "options" bytea NULL, "survey_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "survey_questions"
COMMENT ON TABLE "survey_questions" IS '问卷问题表';
-- set comment to column: "id" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "text" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."text" IS '问题的内容';
-- set comment to column: "type" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."type" IS '问题类型，单选/多选等';
-- set comment to column: "options" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."options" IS '选项内容，单选/多选的内容';
-- set comment to column: "survey_id" on table: "survey_questions"
COMMENT ON COLUMN "survey_questions" ."survey_id" IS '问卷 ID';
-- create "survey_responses" table
CREATE TABLE "survey_responses" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "survey_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "survey_responses"
COMMENT ON TABLE "survey_responses" IS '问卷调查结果储存表';
-- set comment to column: "id" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "survey_id" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."survey_id" IS '问卷 ID';
-- set comment to column: "user_id" on table: "survey_responses"
COMMENT ON COLUMN "survey_responses" ."user_id" IS '用户 ID';
