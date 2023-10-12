-- modify "mission_orders" table
ALTER TABLE "mission_orders" ADD COLUMN "expired_warning_time" timestamptz NULL;
-- set comment to column: "expired_warning_time" on table: "mission_orders"
COMMENT ON COLUMN "mission_orders" ."expired_warning_time" IS '包时任务到期提醒时间（发了通知提醒就更新该时间）';
