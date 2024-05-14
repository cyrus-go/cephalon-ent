-- modify "gpus" table
ALTER TABLE "gpus" ADD COLUMN "trouble_deduct_amount" bigint NOT NULL DEFAULT 0;
-- set comment to column: "trouble_deduct_amount" on table: "gpus"
COMMENT ON COLUMN "gpus" ."trouble_deduct_amount" IS '故障扣费金额，单位：分/小时';
-- create "trouble_deducts" table
CREATE TABLE "trouble_deducts" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "started_at" timestamptz NOT NULL, "finished_at" timestamptz NOT NULL, "time_of_duration" double precision NOT NULL DEFAULT 0, "amount" bigint NOT NULL DEFAULT 0, "status" character varying NOT NULL DEFAULT 'pending', "device_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "trouble_deducts"
COMMENT ON TABLE "trouble_deducts" IS '故障扣费记录，节点故障时，需要扣费，记录到这个表里';
-- set comment to column: "id" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "started_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."started_at" IS '故障开始时刻';
-- set comment to column: "finished_at" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."finished_at" IS '故障结束时刻';
-- set comment to column: "time_of_duration" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."time_of_duration" IS '持续时长，单位：小时';
-- set comment to column: "amount" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."amount" IS '扣费金额，单位：分';
-- set comment to column: "status" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."status" IS '状态';
-- set comment to column: "device_id" on table: "trouble_deducts"
COMMENT ON COLUMN "trouble_deducts" ."device_id" IS '设备 id';
