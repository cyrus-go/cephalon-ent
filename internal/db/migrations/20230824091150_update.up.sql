-- modify "missions" table
ALTER TABLE "missions" ADD COLUMN "gpu_version" character varying NOT NULL DEFAULT 'RTX 2060';
-- set comment to column: "gpu_version" on table: "missions"
COMMENT ON COLUMN "missions" ."gpu_version" IS '最低可接显卡';
