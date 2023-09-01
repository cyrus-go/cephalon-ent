-- modify "prices" table
ALTER TABLE "prices" ADD COLUMN "is_deprecated" boolean NOT NULL DEFAULT false;
-- set comment to column: "is_deprecated" on table: "prices"
COMMENT ON COLUMN "prices" ."is_deprecated" IS '价格是否屏蔽，前端置灰，硬选也可以';
