-- reverse: modify "missions" table
ALTER TABLE "missions" ALTER COLUMN "call_back_info" SET NOT NULL, ALTER COLUMN "call_back_info" SET DEFAULT '';
