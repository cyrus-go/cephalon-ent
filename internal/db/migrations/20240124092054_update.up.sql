-- drop index "cdkinfo_cdk_number" from table: "cdk_infos"
DROP INDEX "cdkinfo_cdk_number";
-- create index "cdkinfo_cdk_number" to table: "cdk_infos"
CREATE UNIQUE INDEX "cdkinfo_cdk_number" ON "cdk_infos" ("cdk_number");
