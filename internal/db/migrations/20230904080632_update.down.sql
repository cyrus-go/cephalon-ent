-- reverse: set comment to column: "second_hmac_key" on table: "missions"
COMMENT ON COLUMN "missions" ."second_hmac_key" IS '';
-- reverse: modify "missions" table
ALTER TABLE "missions" DROP COLUMN "second_hmac_key";
