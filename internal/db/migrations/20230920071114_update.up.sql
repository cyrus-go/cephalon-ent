-- modify "missions" table
ALTER TABLE "missions" ALTER COLUMN "call_back_info" DROP NOT NULL, ALTER COLUMN "call_back_info" DROP DEFAULT;
