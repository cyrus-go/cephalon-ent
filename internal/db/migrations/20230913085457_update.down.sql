-- reverse: set comment to column: "invite_id" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."invite_id" IS '';
-- reverse: set comment to column: "status" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."status" IS '';
-- reverse: set comment to column: "ended_at" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."ended_at" IS '';
-- reverse: set comment to column: "started_at" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."started_at" IS '';
-- reverse: set comment to column: "type" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."type" IS '';
-- reverse: set comment to column: "name" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."name" IS '';
-- reverse: create "campaigns" table
DROP TABLE "campaigns";
-- reverse: set comment to column: "campaign_id" on table: "invites"
COMMENT ON COLUMN "invites" ."campaign_id" IS '';
-- reverse: modify "invites" table
ALTER TABLE "invites" DROP COLUMN "campaign_id";
