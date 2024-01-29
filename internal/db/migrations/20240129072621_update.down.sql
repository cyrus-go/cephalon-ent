-- reverse: set comment to column: "user_id" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."user_id" IS '';
-- reverse: set comment to column: "lotto_id" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."lotto_id" IS '';
-- reverse: set comment to column: "remain_lotto_count" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."remain_lotto_count" IS '';
-- reverse: set comment to column: "deleted_at" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "lotto_user_counts"
COMMENT ON COLUMN "lotto_user_counts" ."id" IS '';
-- reverse: set comment to table: "lotto_user_counts"
COMMENT ON TABLE "lotto_user_counts" IS '';
-- reverse: create "lotto_user_counts" table
DROP TABLE "lotto_user_counts";
-- reverse: set comment to column: "user_id" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."user_id" IS '';
-- reverse: set comment to column: "lotto_prize_id" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."lotto_prize_id" IS '';
-- reverse: set comment to column: "lotto_id" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."lotto_id" IS '';
-- reverse: set comment to column: "remain_lotto_count" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."remain_lotto_count" IS '';
-- reverse: set comment to column: "status" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."status" IS '';
-- reverse: set comment to column: "result" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."result" IS '';
-- reverse: set comment to column: "deleted_at" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "lotto_records"
COMMENT ON COLUMN "lotto_records" ."id" IS '';
-- reverse: set comment to table: "lotto_records"
COMMENT ON TABLE "lotto_records" IS '';
-- reverse: create "lotto_records" table
DROP TABLE "lotto_records";
-- reverse: set comment to column: "lotto_id" on table: "prizes"
COMMENT ON COLUMN "prizes" ."lotto_id" IS '';
-- reverse: set comment to column: "status" on table: "prizes"
COMMENT ON COLUMN "prizes" ."status" IS '';
-- reverse: set comment to column: "name" on table: "prizes"
COMMENT ON COLUMN "prizes" ."name" IS '';
-- reverse: set comment to column: "weight" on table: "prizes"
COMMENT ON COLUMN "prizes" ."weight" IS '';
-- reverse: set comment to column: "level_name" on table: "prizes"
COMMENT ON COLUMN "prizes" ."level_name" IS '';
-- reverse: set comment to column: "deleted_at" on table: "prizes"
COMMENT ON COLUMN "prizes" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "prizes"
COMMENT ON COLUMN "prizes" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "prizes"
COMMENT ON COLUMN "prizes" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "prizes"
COMMENT ON COLUMN "prizes" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "prizes"
COMMENT ON COLUMN "prizes" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "prizes"
COMMENT ON COLUMN "prizes" ."id" IS '';
-- reverse: set comment to table: "prizes"
COMMENT ON TABLE "prizes" IS '';
-- reverse: create index "lottoprize_name" to table: "prizes"
DROP INDEX "lottoprize_name";
-- reverse: create "prizes" table
DROP TABLE "prizes";
-- reverse: set comment to column: "user_id" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."user_id" IS '';
-- reverse: set comment to column: "lotto_id" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."lotto_id" IS '';
-- reverse: set comment to column: "recharge_amount" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."recharge_amount" IS '';
-- reverse: set comment to column: "type" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."type" IS '';
-- reverse: set comment to column: "count" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."count" IS '';
-- reverse: set comment to column: "deleted_at" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "lotto_get_count_records"
COMMENT ON COLUMN "lotto_get_count_records" ."id" IS '';
-- reverse: set comment to table: "lotto_get_count_records"
COMMENT ON TABLE "lotto_get_count_records" IS '';
-- reverse: create "lotto_get_count_records" table
DROP TABLE "lotto_get_count_records";
-- reverse: set comment to column: "status" on table: "lottos"
COMMENT ON COLUMN "lottos" ."status" IS '';
-- reverse: set comment to column: "ended_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."ended_at" IS '';
-- reverse: set comment to column: "started_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."started_at" IS '';
-- reverse: set comment to column: "total_weight" on table: "lottos"
COMMENT ON COLUMN "lottos" ."total_weight" IS '';
-- reverse: set comment to column: "name" on table: "lottos"
COMMENT ON COLUMN "lottos" ."name" IS '';
-- reverse: set comment to column: "deleted_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "lottos"
COMMENT ON COLUMN "lottos" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "lottos"
COMMENT ON COLUMN "lottos" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "lottos"
COMMENT ON COLUMN "lottos" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "lottos"
COMMENT ON COLUMN "lottos" ."id" IS '';
-- reverse: set comment to table: "lottos"
COMMENT ON TABLE "lottos" IS '';
-- reverse: create index "lotto_name" to table: "lottos"
DROP INDEX "lotto_name";
-- reverse: create "lottos" table
DROP TABLE "lottos";
