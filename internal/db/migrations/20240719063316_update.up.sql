-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "use_auth" boolean NOT NULL DEFAULT false, ADD COLUMN "old_mission_id" bigint NULL DEFAULT 0;
-- set comment to column: "use_auth" on table: "missions"
COMMENT ON COLUMN "missions" ."use_auth" IS '是否需要使用鉴权（应用开启后需要账号密码登陆验证）';
-- set comment to column: "old_mission_id" on table: "missions"
COMMENT ON COLUMN "missions" ."old_mission_id" IS '外键，重新开机的旧应用 ID';
