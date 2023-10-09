-- modify "prices" table
ALTER TABLE "prices" ADD COLUMN "renewal_type" character varying NOT NULL DEFAULT 'unknow';
-- set comment to column: "renewal_type" on table: "prices"
COMMENT ON COLUMN "prices" ."renewal_type" IS '包时类型，只有包时任务才有';
