-- modify "extra_service_orders" table
ALTER TABLE "extra_service_orders" ALTER COLUMN "started_at" SET NOT NULL, ALTER COLUMN "started_at" SET DEFAULT CURRENT_TIMESTAMP, ALTER COLUMN "finished_at" SET NOT NULL, ALTER COLUMN "finished_at" SET DEFAULT CURRENT_TIMESTAMP;
