-- create index "missionfailedfeedback_user_id_mission_id" to table: "mission_failed_feedbacks"
CREATE INDEX "missionfailedfeedback_user_id_mission_id" ON "mission_failed_feedbacks" ("user_id", "mission_id");
