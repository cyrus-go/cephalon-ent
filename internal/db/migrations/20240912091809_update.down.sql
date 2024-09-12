-- reverse: set comment to column: "symbol_id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."symbol_id" IS '';
-- reverse: create index "withdrawrecord_transfer_order_id" to table: "withdraw_records"
DROP INDEX "withdrawrecord_transfer_order_id";
-- reverse: create index "withdrawrecord_operate_user_id" to table: "withdraw_records"
DROP INDEX "withdrawrecord_operate_user_id";
-- reverse: modify "withdraw_records" table
ALTER TABLE "withdraw_records" DROP COLUMN "symbol_id";
-- reverse: set comment to column: "channel" on table: "users"
COMMENT ON COLUMN "users" ."channel" IS '';
-- reverse: create index "user_channel" to table: "users"
DROP INDEX "user_channel";
-- reverse: create index "user_applet_parent_id" to table: "users"
DROP INDEX "user_applet_parent_id";
-- reverse: create index "user_parent_id" to table: "users"
DROP INDEX "user_parent_id";
-- reverse: create index "user_email" to table: "users"
DROP INDEX "user_email";
-- reverse: create index "user_phone" to table: "users"
DROP INDEX "user_phone";
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "channel";
-- reverse: set comment to column: "hosting_type" on table: "devices"
COMMENT ON COLUMN "devices" ."hosting_type" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "hosting_type";
