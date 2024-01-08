-- create index "artworklike_user_id" to table: "artwork_likes"
CREATE INDEX "artworklike_user_id" ON "artwork_likes" ("user_id");
-- create index "artworklike_artwork_id" to table: "artwork_likes"
CREATE INDEX "artworklike_artwork_id" ON "artwork_likes" ("artwork_id");
-- create index "artwork_author_id" to table: "artworks"
CREATE INDEX "artwork_author_id" ON "artworks" ("author_id");
-- create index "bill_source_user_id" to table: "bills"
CREATE INDEX "bill_source_user_id" ON "bills" ("source_user_id");
-- create index "bill_target_user_id" to table: "bills"
CREATE INDEX "bill_target_user_id" ON "bills" ("target_user_id");
-- create index "bill_order_id" to table: "bills"
CREATE INDEX "bill_order_id" ON "bills" ("order_id");
-- create index "bill_invite_id" to table: "bills"
CREATE INDEX "bill_invite_id" ON "bills" ("invite_id");
-- create index "bill_symbol_id" to table: "bills"
CREATE INDEX "bill_symbol_id" ON "bills" ("symbol_id");
-- create index "campaignorder_user_id" to table: "campaign_orders"
CREATE INDEX "campaignorder_user_id" ON "campaign_orders" ("user_id");
-- create index "campaignorder_campaign_id" to table: "campaign_orders"
CREATE INDEX "campaignorder_campaign_id" ON "campaign_orders" ("campaign_id");
-- create index "collect_user_id" to table: "collects"
CREATE INDEX "collect_user_id" ON "collects" ("user_id");
-- create index "devicegpumission_device_id" to table: "device_gpu_missions"
CREATE INDEX "devicegpumission_device_id" ON "device_gpu_missions" ("device_id");
-- create index "devicegpumission_gpu_id" to table: "device_gpu_missions"
CREATE INDEX "devicegpumission_gpu_id" ON "device_gpu_missions" ("gpu_id");
-- create index "device_user_id" to table: "devices"
CREATE INDEX "device_user_id" ON "devices" ("user_id");
-- create index "extraserviceorder_mission_id" to table: "extra_service_orders"
CREATE INDEX "extraserviceorder_mission_id" ON "extra_service_orders" ("mission_id");
-- create index "extraserviceorder_mission_order_id" to table: "extra_service_orders"
CREATE INDEX "extraserviceorder_mission_order_id" ON "extra_service_orders" ("mission_order_id");
-- create index "extraserviceorder_symbol_id" to table: "extra_service_orders"
CREATE INDEX "extraserviceorder_symbol_id" ON "extra_service_orders" ("symbol_id");
-- create index "extraserviceorder_mission_batch_id" to table: "extra_service_orders"
CREATE INDEX "extraserviceorder_mission_batch_id" ON "extra_service_orders" ("mission_batch_id");
-- create index "extraserviceprice_extra_service_id" to table: "extra_service_prices"
CREATE INDEX "extraserviceprice_extra_service_id" ON "extra_service_prices" ("extra_service_id");
-- create index "frpcinfo_frps_id" to table: "frpc_infos"
CREATE INDEX "frpcinfo_frps_id" ON "frpc_infos" ("frps_id");
-- create index "frpcinfo_device_id" to table: "frpc_infos"
CREATE INDEX "frpcinfo_device_id" ON "frpc_infos" ("device_id");
-- create index "invite_user_id" to table: "invites"
CREATE INDEX "invite_user_id" ON "invites" ("user_id");
-- create index "invite_campaign_id" to table: "invites"
CREATE INDEX "invite_campaign_id" ON "invites" ("campaign_id");
-- create index "loginrecord_user_id" to table: "login_records"
CREATE INDEX "loginrecord_user_id" ON "login_records" ("user_id");
-- create index "missionbatch_user_id" to table: "mission_batches"
CREATE INDEX "missionbatch_user_id" ON "mission_batches" ("user_id");
-- create index "missionextraservice_mission_id" to table: "mission_extra_services"
CREATE INDEX "missionextraservice_mission_id" ON "mission_extra_services" ("mission_id");
-- create index "missionextraservice_extra_service_id" to table: "mission_extra_services"
CREATE INDEX "missionextraservice_extra_service_id" ON "mission_extra_services" ("extra_service_id");
-- create index "missionorder_consume_user_id" to table: "mission_orders"
CREATE INDEX "missionorder_consume_user_id" ON "mission_orders" ("consume_user_id");
-- create index "missionorder_produce_user_id" to table: "mission_orders"
CREATE INDEX "missionorder_produce_user_id" ON "mission_orders" ("produce_user_id");
-- create index "missionorder_symbol_id" to table: "mission_orders"
CREATE INDEX "missionorder_symbol_id" ON "mission_orders" ("symbol_id");
-- create index "missionorder_mission_batch_id" to table: "mission_orders"
CREATE INDEX "missionorder_mission_batch_id" ON "mission_orders" ("mission_batch_id");
-- create index "missionorder_mission_id" to table: "mission_orders"
CREATE INDEX "missionorder_mission_id" ON "mission_orders" ("mission_id");
-- create index "missionorder_device_id" to table: "mission_orders"
CREATE INDEX "missionorder_device_id" ON "mission_orders" ("device_id");
-- create index "missionproduction_mission_id" to table: "mission_productions"
CREATE INDEX "missionproduction_mission_id" ON "mission_productions" ("mission_id");
-- create index "missionproduction_user_id" to table: "mission_productions"
CREATE INDEX "missionproduction_user_id" ON "mission_productions" ("user_id");
-- create index "missionproduction_device_id" to table: "mission_productions"
CREATE INDEX "missionproduction_device_id" ON "mission_productions" ("device_id");
-- create index "mission_mission_kind_id" to table: "missions"
CREATE INDEX "mission_mission_kind_id" ON "missions" ("mission_kind_id");
-- create index "mission_user_id" to table: "missions"
CREATE INDEX "mission_user_id" ON "missions" ("user_id");
-- create index "mission_mission_batch_id" to table: "missions"
CREATE INDEX "mission_mission_batch_id" ON "missions" ("mission_batch_id");
-- set comment to table: "missions"
COMMENT ON TABLE "missions" IS '任务，具备任务要求，记录完成情况和结果，金额相关信息在 mission_orders 等订单侧';
-- create index "price_gpu_id" to table: "prices"
CREATE INDEX "price_gpu_id" ON "prices" ("gpu_id");
-- create index "profitsetting_user_id" to table: "profit_settings"
CREATE INDEX "profitsetting_user_id" ON "profit_settings" ("user_id");
-- create index "rechargeorder_user_id" to table: "recharge_orders"
CREATE INDEX "rechargeorder_user_id" ON "recharge_orders" ("user_id");
-- create index "rechargeorder_social_id" to table: "recharge_orders"
CREATE INDEX "rechargeorder_social_id" ON "recharge_orders" ("social_id");
-- create index "rechargeorder_campaign_order_id" to table: "recharge_orders"
CREATE INDEX "rechargeorder_campaign_order_id" ON "recharge_orders" ("campaign_order_id");
-- create index "renewalagreement_user_id" to table: "renewal_agreements"
CREATE INDEX "renewalagreement_user_id" ON "renewal_agreements" ("user_id");
-- create index "renewalagreement_mission_id" to table: "renewal_agreements"
CREATE INDEX "renewalagreement_mission_id" ON "renewal_agreements" ("mission_id");
-- create index "transferorder_source_user_id" to table: "transfer_orders"
CREATE INDEX "transferorder_source_user_id" ON "transfer_orders" ("source_user_id");
-- create index "transferorder_target_user_id" to table: "transfer_orders"
CREATE INDEX "transferorder_target_user_id" ON "transfer_orders" ("target_user_id");
-- create index "transferorder_social_id" to table: "transfer_orders"
CREATE INDEX "transferorder_social_id" ON "transfer_orders" ("social_id");
-- create index "transferorder_symbol_id" to table: "transfer_orders"
CREATE INDEX "transferorder_symbol_id" ON "transfer_orders" ("symbol_id");
-- create index "userdevice_user_id" to table: "user_devices"
CREATE INDEX "userdevice_user_id" ON "user_devices" ("user_id");
-- create index "userdevice_device_id" to table: "user_devices"
CREATE INDEX "userdevice_device_id" ON "user_devices" ("device_id");
-- create index "vxaccount_user_id" to table: "vx_accounts"
CREATE INDEX "vxaccount_user_id" ON "vx_accounts" ("user_id");
-- create index "vxsocial_user_id" to table: "vx_socials"
CREATE INDEX "vxsocial_user_id" ON "vx_socials" ("user_id");
-- create index "wallet_user_id" to table: "wallets"
CREATE INDEX "wallet_user_id" ON "wallets" ("user_id");
-- create index "wallet_symbol_id" to table: "wallets"
CREATE INDEX "wallet_symbol_id" ON "wallets" ("symbol_id");
-- create index "withdrawaccount_user_id" to table: "withdraw_accounts"
CREATE INDEX "withdrawaccount_user_id" ON "withdraw_accounts" ("user_id");
