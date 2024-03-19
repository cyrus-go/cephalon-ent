-- modify "transfer_orders" table
ALTER TABLE "transfer_orders" ADD COLUMN "withdraw_account" character varying NOT NULL DEFAULT '';
-- set comment to column: "withdraw_account" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."withdraw_account" IS '提现账户（类型为提现才有数据）';
