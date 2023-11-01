-- reverse: drop index "renewal_agreements_mission_id_key" from table: "renewal_agreements"
CREATE UNIQUE INDEX "renewal_agreements_mission_id_key" ON "renewal_agreements" ("mission_id");
