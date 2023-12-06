-- reverse: modify "missions" table
ALTER TABLE "missions" ALTER COLUMN "free_at" SET DEFAULT '0001-01-01 08:05:43+08:05:43';
-- reverse: modify "mission_orders" table
ALTER TABLE "mission_orders" ALTER COLUMN "lately_settled_at" SET DEFAULT '0001-01-01 08:05:43+08:05:43';
