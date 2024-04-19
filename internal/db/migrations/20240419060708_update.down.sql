-- reverse: set comment to column: "reject_reason" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."reject_reason" IS '';
-- reverse: set comment to column: "operate_user_id" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."operate_user_id" IS '操作的用户 id，手动充值才有数据，默认为 0';
-- reverse: modify "transfer_orders" table
ALTER TABLE "transfer_orders" DROP COLUMN "reject_reason";
