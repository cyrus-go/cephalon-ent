-- reverse: set comment to column: "is_gift_recharge" on table: "surveys"
COMMENT ON COLUMN "surveys" ."is_gift_recharge" IS '';
-- reverse: modify "surveys" table
ALTER TABLE "surveys" DROP COLUMN "is_gift_recharge";
