-- modify "prices" table
ALTER TABLE "prices" ADD COLUMN "is_sensitive" boolean NOT NULL DEFAULT false;
-- set comment to column: "is_sensitive" on table: "prices"
COMMENT ON COLUMN "prices" ."is_sensitive" IS '价格是否敏感，用于特殊类型任务，不能让外部看到选项';
