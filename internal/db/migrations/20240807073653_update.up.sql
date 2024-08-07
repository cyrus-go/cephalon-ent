-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "rank" character varying NOT NULL DEFAULT 'normal';
-- set comment to column: "rank" on table: "devices"
COMMENT ON COLUMN "devices" ."rank" IS '设备等级(目前阶段就是黑名单)';
