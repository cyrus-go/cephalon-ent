-- reverse: set comment to column: "withdraw_retain_amount" on table: "gpus"
COMMENT ON COLUMN "gpus" ."withdraw_retain_amount" IS '';
-- reverse: modify "gpus" table
ALTER TABLE "gpus" DROP COLUMN "withdraw_retain_amount";
-- reverse: modify "devices" table
ALTER TABLE "devices" ALTER COLUMN "status" SET DEFAULT 'online';
