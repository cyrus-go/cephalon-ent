-- reverse: set comment to column: "operate_user_id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."operate_user_id" IS '';
-- reverse: set comment to column: "user_id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."user_id" IS '';
-- reverse: set comment to column: "transfer_order_id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."transfer_order_id" IS '';
-- reverse: set comment to column: "reject_reason" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."reject_reason" IS '';
-- reverse: set comment to column: "status" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."status" IS '';
-- reverse: set comment to column: "real_amount" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."real_amount" IS '';
-- reverse: set comment to column: "rate" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."rate" IS '';
-- reverse: set comment to column: "remain_amount" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."remain_amount" IS '';
-- reverse: set comment to column: "amount" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."amount" IS '';
-- reverse: set comment to column: "type" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."type" IS '';
-- reverse: set comment to column: "withdraw_account" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."withdraw_account" IS '';
-- reverse: set comment to column: "deleted_at" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."id" IS '';
-- reverse: set comment to table: "withdraw_records"
COMMENT ON TABLE "withdraw_records" IS '';
-- reverse: create index "withdrawrecord_user_id" to table: "withdraw_records"
DROP INDEX "withdrawrecord_user_id";
-- reverse: create index "withdraw_records_transfer_order_id_key" to table: "withdraw_records"
DROP INDEX "withdraw_records_transfer_order_id_key";
-- reverse: create "withdraw_records" table
DROP TABLE "withdraw_records";
