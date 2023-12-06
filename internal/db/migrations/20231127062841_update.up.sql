-- modify "mission_orders" table
ALTER TABLE "mission_orders" ALTER COLUMN "lately_settled_at" SET DEFAULT CURRENT_TIMESTAMP;
-- modify "missions" table
ALTER TABLE "missions" ALTER COLUMN "free_at" SET DEFAULT CURRENT_TIMESTAMP;
