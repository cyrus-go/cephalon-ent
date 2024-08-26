-- reverse: set comment to column: "gap_random_min" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."gap_random_min" IS '';
-- reverse: set comment to column: "gap_random_max" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."gap_random_max" IS '';
-- reverse: set comment to column: "gap_base" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."gap_base" IS '';
-- reverse: set comment to column: "gpu_version" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."gpu_version" IS '';
-- reverse: set comment to column: "stability_level" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."stability_level" IS '';
-- reverse: set comment to column: "deleted_at" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "gift_mission_configs"
COMMENT ON COLUMN "gift_mission_configs" ."id" IS '';
-- reverse: set comment to table: "gift_mission_configs"
COMMENT ON TABLE "gift_mission_configs" IS '';
-- reverse: create index "giftmissionconfig_stability_level_gpu_version" to table: "gift_mission_configs"
DROP INDEX "giftmissionconfig_stability_level_gpu_version";
-- reverse: create "gift_mission_configs" table
DROP TABLE "gift_mission_configs";
-- reverse: set comment to column: "gift_mission_config_id" on table: "devices"
COMMENT ON COLUMN "devices" ."gift_mission_config_id" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "gift_mission_config_id";
