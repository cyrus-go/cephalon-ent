-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "hosting_type" character varying NOT NULL DEFAULT 'no';
-- set comment to column: "hosting_type" on table: "devices"
COMMENT ON COLUMN "devices" ."hosting_type" IS '托管类型，非托管/半托管/全托管等';
-- modify "users" table
ALTER TABLE "users" ADD COLUMN "channel" character varying NOT NULL DEFAULT 'no_channel';
-- create index "user_phone" to table: "users"
CREATE INDEX "user_phone" ON "users" ("phone");
-- create index "user_email" to table: "users"
CREATE INDEX "user_email" ON "users" ("email");
-- create index "user_parent_id" to table: "users"
CREATE INDEX "user_parent_id" ON "users" ("parent_id");
-- create index "user_applet_parent_id" to table: "users"
CREATE INDEX "user_applet_parent_id" ON "users" ("applet_parent_id");
-- create index "user_channel" to table: "users"
CREATE INDEX "user_channel" ON "users" ("channel");
-- set comment to column: "channel" on table: "users"
COMMENT ON COLUMN "users" ."channel" IS '渠道身份，默认为非渠道用户';
-- modify "withdraw_records" table
ALTER TABLE "withdraw_records" ADD COLUMN "symbol_id" bigint NOT NULL DEFAULT 0;
-- create index "withdrawrecord_operate_user_id" to table: "withdraw_records"
CREATE INDEX "withdrawrecord_operate_user_id" ON "withdraw_records" ("operate_user_id");
-- create index "withdrawrecord_transfer_order_id" to table: "withdraw_records"
CREATE INDEX "withdrawrecord_transfer_order_id" ON "withdraw_records" ("transfer_order_id");
-- set comment to column: "symbol_id" on table: "withdraw_records"
COMMENT ON COLUMN "withdraw_records" ."symbol_id" IS '消费的外键币种 id';
