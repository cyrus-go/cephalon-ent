-- modify "gpus" table
ALTER TABLE "gpus" ADD COLUMN "power" bigint NOT NULL DEFAULT 0;
-- set comment to column: "power" on table: "gpus"
COMMENT ON COLUMN "gpus" ."power" IS '显卡能力值';
