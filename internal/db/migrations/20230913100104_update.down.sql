-- reverse: set comment to column: "is_recharge" on table: "users"
COMMENT ON COLUMN "users" ."is_recharge" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "is_recharge";
