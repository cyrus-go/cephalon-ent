-- reverse: set comment to column: "channel_ratio" on table: "users"
COMMENT ON COLUMN "users" ."channel_ratio" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "channel_ratio";
