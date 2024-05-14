-- reverse: set comment to column: "device_id" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."device_id" IS '';
-- reverse: set comment to column: "status" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."status" IS '';
-- reverse: set comment to column: "amount" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."amount" IS '';
-- reverse: set comment to column: "time_of_duration" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."time_of_duration" IS '';
-- reverse: set comment to column: "finished_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."finished_at" IS '';
-- reverse: set comment to column: "started_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."started_at" IS '';
-- reverse: set comment to column: "deleted_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."id" IS '';
-- reverse: set comment to table: "trouble_deducts"
COMMENT ON TABLE "trouble_deducts" IS '';
-- reverse: create "trouble_deducts" table
DROP TABLE "trouble_deducts";
-- reverse: set comment to column: "trouble_deduct_amount" on table: "gpus"
COMMENT ON COLUMN "gpus" ."trouble_deduct_amount" IS '';
-- reverse: modify "gpus" table
ALTER TABLE "gpus" DROP COLUMN "trouble_deduct_amount";
