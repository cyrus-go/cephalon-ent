-- create "prices" table
CREATE TABLE "prices" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "gpu_version" character varying NOT NULL DEFAULT 'RTX 2060', "mission_type" character varying NOT NULL DEFAULT 'txt2img', "mission_category" character varying NOT NULL DEFAULT 'SD', "mission_billing_type" character varying NOT NULL DEFAULT 'count', "cep" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "gpu_version" on table: "prices"
COMMENT ON COLUMN "prices" ."gpu_version" IS '显卡型号';
-- set comment to column: "mission_type" on table: "prices"
COMMENT ON COLUMN "prices" ."mission_type" IS '任务类型';
-- set comment to column: "mission_category" on table: "prices"
COMMENT ON COLUMN "prices" ."mission_category" IS '任务大类';
-- set comment to column: "mission_billing_type" on table: "prices"
COMMENT ON COLUMN "prices" ."mission_billing_type" IS '任务计费类型';
-- set comment to column: "cep" on table: "prices"
COMMENT ON COLUMN "prices" ."cep" IS '任务单价';
