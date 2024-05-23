-- set comment to column: "operate_user_id" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."operate_user_id" IS '操作的用户 id，管理后台手动充值才有数据，默认为 0';
