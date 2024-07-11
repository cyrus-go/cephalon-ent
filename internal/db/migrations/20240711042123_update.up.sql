-- modify "surveys" table
ALTER TABLE "surveys" ADD COLUMN "is_gift_recharge" boolean NOT NULL DEFAULT false;
-- set comment to column: "is_gift_recharge" on table: "surveys"
COMMENT ON COLUMN "surveys" ."is_gift_recharge" IS '问卷是否参加充值活动';
