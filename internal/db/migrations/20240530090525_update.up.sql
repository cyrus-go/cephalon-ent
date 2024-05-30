-- modify "income_manages" table
ALTER TABLE "income_manages" ADD COLUMN "last_edited_user_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "last_edited_user_id" on table: "income_manages"
COMMENT ON COLUMN "income_manages" ."last_edited_user_id" IS '审批前最后一次编辑的用戶 id';
