-- modify "users" table
ALTER TABLE "users" ADD COLUMN "channel_ratio" bigint NOT NULL DEFAULT 0;
-- set comment to column: "channel_ratio" on table: "users"
COMMENT ON COLUMN "users" ."channel_ratio" IS '渠道分成比例';
