-- modify "missions" table
ALTER TABLE "missions" ALTER COLUMN "gpu_num" TYPE bigint;
-- modify "users" table
ALTER TABLE "users" ALTER COLUMN "github_id" TYPE character varying, ALTER COLUMN "github_id" SET DEFAULT '';
