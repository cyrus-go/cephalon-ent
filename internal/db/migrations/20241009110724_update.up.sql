-- modify "invites" table
ALTER TABLE "invites" ADD COLUMN "channel_ratio" bigint NOT NULL DEFAULT 0;
-- set comment to column: "channel_ratio" on table: "invites"
COMMENT ON COLUMN "invites" ."channel_ratio" IS '渠道分成比例';
