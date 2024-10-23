-- reverse: set comment to column: "last_abnormal_at" on table: "devices"
COMMENT ON COLUMN "devices" ."last_abnormal_at" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "last_abnormal_at";
