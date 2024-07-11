-- modify "transfer_orders" table
ALTER TABLE "transfer_orders" ADD COLUMN "gift_type" character varying NOT NULL DEFAULT 'no';
-- set comment to column: "gift_type" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."gift_type" IS '充值订单活动赠送的类型';
