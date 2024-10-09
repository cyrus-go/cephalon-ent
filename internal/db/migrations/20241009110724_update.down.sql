-- reverse: set comment to column: "channel_ratio" on table: "invites"
COMMENT ON COLUMN "invites" ."channel_ratio" IS '';
-- reverse: modify "invites" table
ALTER TABLE "invites" DROP COLUMN "channel_ratio";
