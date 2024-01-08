-- reverse: create index "withdrawaccount_user_id" to table: "withdraw_accounts"
DROP INDEX "withdrawaccount_user_id";
-- reverse: create index "wallet_symbol_id" to table: "wallets"
DROP INDEX "wallet_symbol_id";
-- reverse: create index "wallet_user_id" to table: "wallets"
DROP INDEX "wallet_user_id";
-- reverse: create index "vxsocial_user_id" to table: "vx_socials"
DROP INDEX "vxsocial_user_id";
-- reverse: create index "vxaccount_user_id" to table: "vx_accounts"
DROP INDEX "vxaccount_user_id";
-- reverse: create index "userdevice_device_id" to table: "user_devices"
DROP INDEX "userdevice_device_id";
-- reverse: create index "userdevice_user_id" to table: "user_devices"
DROP INDEX "userdevice_user_id";
-- reverse: create index "transferorder_symbol_id" to table: "transfer_orders"
DROP INDEX "transferorder_symbol_id";
-- reverse: create index "transferorder_social_id" to table: "transfer_orders"
DROP INDEX "transferorder_social_id";
-- reverse: create index "transferorder_target_user_id" to table: "transfer_orders"
DROP INDEX "transferorder_target_user_id";
-- reverse: create index "transferorder_source_user_id" to table: "transfer_orders"
DROP INDEX "transferorder_source_user_id";
-- reverse: create index "renewalagreement_mission_id" to table: "renewal_agreements"
DROP INDEX "renewalagreement_mission_id";
-- reverse: create index "renewalagreement_user_id" to table: "renewal_agreements"
DROP INDEX "renewalagreement_user_id";
-- reverse: create index "rechargeorder_campaign_order_id" to table: "recharge_orders"
DROP INDEX "rechargeorder_campaign_order_id";
-- reverse: create index "rechargeorder_social_id" to table: "recharge_orders"
DROP INDEX "rechargeorder_social_id";
-- reverse: create index "rechargeorder_user_id" to table: "recharge_orders"
DROP INDEX "rechargeorder_user_id";
-- reverse: create index "profitsetting_user_id" to table: "profit_settings"
DROP INDEX "profitsetting_user_id";
-- reverse: create index "price_gpu_id" to table: "prices"
DROP INDEX "price_gpu_id";
-- reverse: set comment to table: "missions"
COMMENT ON TABLE "missions" IS '任务，具备任务要求，记录完成情况和结果，金额相关信息在 mission_consume_orders 等订单侧';
-- reverse: create index "mission_mission_batch_id" to table: "missions"
DROP INDEX "mission_mission_batch_id";
-- reverse: create index "mission_user_id" to table: "missions"
DROP INDEX "mission_user_id";
-- reverse: create index "mission_mission_kind_id" to table: "missions"
DROP INDEX "mission_mission_kind_id";
-- reverse: create index "missionproduction_device_id" to table: "mission_productions"
DROP INDEX "missionproduction_device_id";
-- reverse: create index "missionproduction_user_id" to table: "mission_productions"
DROP INDEX "missionproduction_user_id";
-- reverse: create index "missionproduction_mission_id" to table: "mission_productions"
DROP INDEX "missionproduction_mission_id";
-- reverse: create index "missionorder_device_id" to table: "mission_orders"
DROP INDEX "missionorder_device_id";
-- reverse: create index "missionorder_mission_id" to table: "mission_orders"
DROP INDEX "missionorder_mission_id";
-- reverse: create index "missionorder_mission_batch_id" to table: "mission_orders"
DROP INDEX "missionorder_mission_batch_id";
-- reverse: create index "missionorder_symbol_id" to table: "mission_orders"
DROP INDEX "missionorder_symbol_id";
-- reverse: create index "missionorder_produce_user_id" to table: "mission_orders"
DROP INDEX "missionorder_produce_user_id";
-- reverse: create index "missionorder_consume_user_id" to table: "mission_orders"
DROP INDEX "missionorder_consume_user_id";
-- reverse: create index "missionextraservice_extra_service_id" to table: "mission_extra_services"
DROP INDEX "missionextraservice_extra_service_id";
-- reverse: create index "missionextraservice_mission_id" to table: "mission_extra_services"
DROP INDEX "missionextraservice_mission_id";
-- reverse: create index "missionbatch_user_id" to table: "mission_batches"
DROP INDEX "missionbatch_user_id";
-- reverse: create index "loginrecord_user_id" to table: "login_records"
DROP INDEX "loginrecord_user_id";
-- reverse: create index "invite_campaign_id" to table: "invites"
DROP INDEX "invite_campaign_id";
-- reverse: create index "invite_user_id" to table: "invites"
DROP INDEX "invite_user_id";
-- reverse: create index "frpcinfo_device_id" to table: "frpc_infos"
DROP INDEX "frpcinfo_device_id";
-- reverse: create index "frpcinfo_frps_id" to table: "frpc_infos"
DROP INDEX "frpcinfo_frps_id";
-- reverse: create index "extraserviceprice_extra_service_id" to table: "extra_service_prices"
DROP INDEX "extraserviceprice_extra_service_id";
-- reverse: create index "extraserviceorder_mission_batch_id" to table: "extra_service_orders"
DROP INDEX "extraserviceorder_mission_batch_id";
-- reverse: create index "extraserviceorder_symbol_id" to table: "extra_service_orders"
DROP INDEX "extraserviceorder_symbol_id";
-- reverse: create index "extraserviceorder_mission_order_id" to table: "extra_service_orders"
DROP INDEX "extraserviceorder_mission_order_id";
-- reverse: create index "extraserviceorder_mission_id" to table: "extra_service_orders"
DROP INDEX "extraserviceorder_mission_id";
-- reverse: create index "device_user_id" to table: "devices"
DROP INDEX "device_user_id";
-- reverse: create index "devicegpumission_gpu_id" to table: "device_gpu_missions"
DROP INDEX "devicegpumission_gpu_id";
-- reverse: create index "devicegpumission_device_id" to table: "device_gpu_missions"
DROP INDEX "devicegpumission_device_id";
-- reverse: create index "collect_user_id" to table: "collects"
DROP INDEX "collect_user_id";
-- reverse: create index "campaignorder_campaign_id" to table: "campaign_orders"
DROP INDEX "campaignorder_campaign_id";
-- reverse: create index "campaignorder_user_id" to table: "campaign_orders"
DROP INDEX "campaignorder_user_id";
-- reverse: create index "bill_symbol_id" to table: "bills"
DROP INDEX "bill_symbol_id";
-- reverse: create index "bill_invite_id" to table: "bills"
DROP INDEX "bill_invite_id";
-- reverse: create index "bill_order_id" to table: "bills"
DROP INDEX "bill_order_id";
-- reverse: create index "bill_target_user_id" to table: "bills"
DROP INDEX "bill_target_user_id";
-- reverse: create index "bill_source_user_id" to table: "bills"
DROP INDEX "bill_source_user_id";
-- reverse: create index "artwork_author_id" to table: "artworks"
DROP INDEX "artwork_author_id";
-- reverse: create index "artworklike_artwork_id" to table: "artwork_likes"
DROP INDEX "artworklike_artwork_id";
-- reverse: create index "artworklike_user_id" to table: "artwork_likes"
DROP INDEX "artworklike_user_id";
