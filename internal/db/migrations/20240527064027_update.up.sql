-- modify "trouble_deducts" table
ALTER TABLE "trouble_deducts" ADD COLUMN "user_id" bigint NOT NULL DEFAULT 0;
-- create index "troublededuct_user_id" to table: "trouble_deducts"
CREATE INDEX "troublededuct_user_id" ON "trouble_deducts" ("user_id");
-- create index "troublededuct_device_id" to table: "trouble_deducts"
CREATE INDEX "troublededuct_device_id" ON "trouble_deducts" ("device_id");
-- set comment to column: "user_id" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."user_id" IS '用戶 id';
