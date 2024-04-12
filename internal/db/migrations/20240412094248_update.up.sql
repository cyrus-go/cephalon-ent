-- modify "transfer_orders" table
ALTER TABLE "transfer_orders" ADD COLUMN "withdraw_rate" bigint NOT NULL DEFAULT 7, ADD COLUMN "withdraw_real_amount" bigint NOT NULL DEFAULT 0;
-- set comment to column: "withdraw_rate" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."withdraw_rate" IS '提现手续费率，100 为基准，比如手续费 7%，值就应该为 7，最大值不能超过 100, 默认 7%';
-- set comment to column: "withdraw_real_amount" on table: "transfer_orders"
COMMENT ON COLUMN "transfer_orders" ."withdraw_real_amount" IS '提现实际到账，单位：cep';
