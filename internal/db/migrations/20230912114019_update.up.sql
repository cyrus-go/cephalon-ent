-- create "invites" table
CREATE TABLE "invites" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "invite_code" character varying NOT NULL DEFAULT '', "share_cep" bigint NOT NULL DEFAULT 0, "reg_cep" bigint NOT NULL DEFAULT 0, "type" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "invite_invite_code" to table: "invites"
CREATE INDEX "invite_invite_code" ON "invites" ("invite_code");
-- set comment to column: "invite_code" on table: "invites"
COMMENT ON COLUMN "invites" ."invite_code" IS '邀请码';
-- set comment to column: "share_cep" on table: "invites"
COMMENT ON COLUMN "invites" ."share_cep" IS '通过此邀请码分享能获得的收益';
-- set comment to column: "reg_cep" on table: "invites"
COMMENT ON COLUMN "invites" ."reg_cep" IS '通过此邀请码注册能获得的收益';
-- set comment to column: "type" on table: "invites"
COMMENT ON COLUMN "invites" ."type" IS '邀请码类型（可以用来区分不同的活动）';
-- set comment to column: "user_id" on table: "invites"
COMMENT ON COLUMN "invites" ."user_id" IS '外键用户 id';
