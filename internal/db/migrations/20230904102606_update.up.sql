-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "inner_uri" character varying NOT NULL DEFAULT '';
-- set comment to column: "inner_uri" on table: "missions"
COMMENT ON COLUMN "missions" ."inner_uri" IS '当 type 为 sd_api 时使用，为转发的 sd 内部接口相对路径';
