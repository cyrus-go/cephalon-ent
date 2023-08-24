-- reverse: set comment to column: "cep" on table: "prices"
COMMENT ON COLUMN "prices" ."cep" IS '';
-- reverse: set comment to column: "mission_billing_type" on table: "prices"
COMMENT ON COLUMN "prices" ."mission_billing_type" IS '';
-- reverse: set comment to column: "mission_category" on table: "prices"
COMMENT ON COLUMN "prices" ."mission_category" IS '';
-- reverse: set comment to column: "mission_type" on table: "prices"
COMMENT ON COLUMN "prices" ."mission_type" IS '';
-- reverse: set comment to column: "gpu_version" on table: "prices"
COMMENT ON COLUMN "prices" ."gpu_version" IS '';
-- reverse: create "prices" table
DROP TABLE "prices";
