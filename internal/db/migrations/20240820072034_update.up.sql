-- modify "devices" table
ALTER TABLE "devices" ADD COLUMN "stability_at" timestamptz NULL;
-- set comment to column: "stability_at" on table: "devices"
COMMENT ON COLUMN "devices" ."stability_at" IS '判定稳定性的时刻，带时区';
