-- reverse: set comment to column: "operate_user_id" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."operate_user_id" IS '操作的用户 id，手动充值或者提现审批才有数据，默认为 0';
