-- reverse: set comment to column: "user_id" on table: "invites"
COMMENT ON COLUMN "invites" ."user_id" IS '';
-- reverse: set comment to column: "type" on table: "invites"
COMMENT ON COLUMN "invites" ."type" IS '';
-- reverse: set comment to column: "reg_cep" on table: "invites"
COMMENT ON COLUMN "invites" ."reg_cep" IS '';
-- reverse: set comment to column: "share_cep" on table: "invites"
COMMENT ON COLUMN "invites" ."share_cep" IS '';
-- reverse: set comment to column: "invite_code" on table: "invites"
COMMENT ON COLUMN "invites" ."invite_code" IS '';
-- reverse: create index "invite_invite_code" to table: "invites"
DROP INDEX "invite_invite_code";
-- reverse: create "invites" table
DROP TABLE "invites";
