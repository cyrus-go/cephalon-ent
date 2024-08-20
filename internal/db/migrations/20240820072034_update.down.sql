-- reverse: set comment to column: "stability_at" on table: "devices"
COMMENT ON COLUMN "devices" ."stability_at" IS '';
-- reverse: modify "devices" table
ALTER TABLE "devices" DROP COLUMN "stability_at";
