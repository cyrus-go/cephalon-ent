-- reverse: set comment to column: "user_id" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."user_id" IS '';
-- reverse: create index "troublededuct_device_id" to table: "trouble_deducts"
DROP INDEX "troublededuct_device_id";
-- reverse: create index "troublededuct_user_id" to table: "trouble_deducts"
DROP INDEX "troublededuct_user_id";
-- reverse: modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" DROP COLUMN "user_id";
