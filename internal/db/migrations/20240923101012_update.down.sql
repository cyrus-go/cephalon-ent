-- reverse: modify "users" table
ALTER TABLE "users" ALTER COLUMN "github_id" TYPE bigint, ALTER COLUMN "github_id" SET DEFAULT 0;
-- reverse: modify "missions" table
ALTER TABLE "missions" ALTER COLUMN "gpu_num" TYPE integer;
