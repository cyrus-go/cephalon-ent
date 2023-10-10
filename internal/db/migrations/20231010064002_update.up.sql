-- modify "renewal_agreements" table
ALTER TABLE "renewal_agreements" ADD COLUMN "last_warning_time" timestamptz NOT NULL;
-- set comment to column: "last_warning_time" on table: "renewal_agreements"
COMMENT ON COLUMN "renewal_agreements" ."last_warning_time" IS '最后一次预警时间';
