-- reverse: set comment to column: "approve_user_id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."approve_user_id" IS '';
-- reverse: set comment to column: "user_id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."user_id" IS '';
-- reverse: set comment to column: "status" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."status" IS '';
-- reverse: set comment to column: "reject_reason" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."reject_reason" IS '';
-- reverse: set comment to column: "last_edited_at" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."last_edited_at" IS '';
-- reverse: set comment to column: "current_balance" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."current_balance" IS '';
-- reverse: set comment to column: "reason" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."reason" IS '';
-- reverse: set comment to column: "amount" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."amount" IS '';
-- reverse: set comment to column: "type" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."type" IS '';
-- reverse: set comment to column: "phone" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."phone" IS '';
-- reverse: set comment to column: "deleted_at" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."id" IS '';
-- reverse: set comment to table: "income_manages"
COMMENT ON TABLE "income_manages" IS '';
-- reverse: create index "incomemanage_approve_user_id" to table: "income_manages"
DROP INDEX "incomemanage_approve_user_id";
-- reverse: create index "incomemanage_user_id" to table: "income_manages"
DROP INDEX "incomemanage_user_id";
-- reverse: create "income_manages" table
DROP TABLE "income_manages";
