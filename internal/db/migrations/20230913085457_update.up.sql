-- modify "invites" table
ALTER TABLE "invites" ADD COLUMN "campaign_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "campaign_id" on table: "invites"
COMMENT ON COLUMN "invites" ."campaign_id" IS '外键活动 id';
-- create "campaigns" table
CREATE TABLE "campaigns" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '', "type" character varying NOT NULL DEFAULT '', "started_at" timestamptz NOT NULL, "ended_at" timestamptz NOT NULL, "status" bigint NOT NULL DEFAULT 0, "invite_id" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- set comment to column: "name" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."name" IS '活动名称';
-- set comment to column: "type" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."type" IS '活动类型';
-- set comment to column: "started_at" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."started_at" IS '活动开始时间';
-- set comment to column: "ended_at" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."ended_at" IS '活动结束时间';
-- set comment to column: "status" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."status" IS '活动状态';
-- set comment to column: "invite_id" on table: "campaigns"
COMMENT ON COLUMN "campaigns" ."invite_id" IS '外键邀请码 id';
