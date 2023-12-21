-- reverse: modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" ALTER COLUMN "finished_at" DROP NOT NULL, ALTER COLUMN "finished_at" DROP DEFAULT, ALTER COLUMN "started_at" DROP NOT NULL, ALTER COLUMN "started_at" DROP DEFAULT;
