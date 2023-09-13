-- modify "users" table
ALTER TABLE "users" ADD COLUMN "is_recharge" boolean NOT NULL DEFAULT false;
-- set comment to column: "is_recharge" on table: "users"
COMMENT ON COLUMN "users" ."is_recharge" IS '是否充值过';
