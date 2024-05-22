-- reverse: set comment to column: "deduct_standard" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."deduct_standard" IS '扣费标准，单位：分';
-- reverse: set comment to column: "amount" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."amount" IS '扣费金额，单位：分';
-- reverse: set comment to column: "trouble_deduct_amount" on table: "gpus"
COMMENT ON COLUMN "gpus" ."trouble_deduct_amount" IS '故障扣费金额，单位：分/小时';
