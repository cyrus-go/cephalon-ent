-- modify "users" table
ALTER TABLE "users" ADD COLUMN "applet_parent_id" bigint NULL DEFAULT 0;
-- set comment to column: "applet_parent_id" on table: "users"
COMMENT ON COLUMN "users" ."applet_parent_id" IS '小程序邀请人用户 id';
