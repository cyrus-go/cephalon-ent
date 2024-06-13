-- modify "devices" table
ALTER TABLE "devices" ALTER COLUMN "status" SET DEFAULT 'offline';
-- modify "gpus" table
ALTER TABLE "gpus" ADD COLUMN "withdraw_retain_amount" bigint NOT NULL DEFAULT 0;
-- set comment to column: "withdraw_retain_amount" on table: "gpus"
COMMENT ON COLUMN "gpus" ."withdraw_retain_amount" IS '提现保留最低金额（押金），单位：厘';
