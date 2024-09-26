-- create "gpu_peaks" table
CREATE TABLE "gpu_peaks" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "version" character varying NOT NULL DEFAULT 'RTX2060', "peak" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "gpu_peaks"
COMMENT ON COLUMN "gpu_peaks" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "gpu_peaks"
COMMENT ON COLUMN "gpu_peaks" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "gpu_peaks"
COMMENT ON COLUMN "gpu_peaks" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "gpu_peaks"
COMMENT ON COLUMN "gpu_peaks" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "gpu_peaks"
COMMENT ON COLUMN "gpu_peaks" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "gpu_peaks"
COMMENT ON COLUMN "gpu_peaks" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "version" on table: "gpu_peaks"
COMMENT ON COLUMN "gpu_peaks" ."version" IS '显卡型号';
-- set comment to column: "peak" on table: "gpu_peaks"
COMMENT ON COLUMN "gpu_peaks" ."peak" IS '每日同时占用峰值';
