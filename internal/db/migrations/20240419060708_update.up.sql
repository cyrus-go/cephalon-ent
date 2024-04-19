-- modify "transfer_orders" table
ALTER TABLE "transfer_orders" ADD COLUMN "reject_reason" character varying NOT NULL DEFAULT '';
-- set comment to column: "operate_user_id" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."operate_user_id" IS '操作的用户 id，手动充值或者提现审批才有数据，默认为 0';
-- set comment to column: "reject_reason" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."reject_reason" IS '提现审批拒绝的理由';
