-- modify "users" table
ALTER TABLE "users" ADD COLUMN "area_code" character varying NOT NULL DEFAULT '+86', ADD COLUMN "email" character varying NOT NULL DEFAULT '';
-- set comment to column: "area_code" on table: "users"
COMMENT ON COLUMN "users" ."area_code" IS '国家区号';
-- set comment to column: "email" on table: "users"
COMMENT ON COLUMN "users" ."email" IS '邮箱';
