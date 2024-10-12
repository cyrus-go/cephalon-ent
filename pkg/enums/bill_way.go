package enums

type BillWay string

const (
	BillWayUnknown                        BillWay = "unknown"
	BillWayRechargeWechat                 BillWay = "recharge_wechat"
	BillWayRechargeAlipay                 BillWay = "recharge_alipay"
	BillWayRechargeTron                   BillWay = "recharge_tron" // TRON 链支付
	BillWayMissionTime                    BillWay = "mission_time"
	BillWayMissionTimePlanHour            BillWay = "mission_time_plan_hour"
	BillWayMissionTimePlanDay             BillWay = "mission_time_plan_day"
	BillWayMissionTimePlanWeek            BillWay = "mission_time_plan_week"
	BillWayMissionTimePlanMonth           BillWay = "mission_time_plan_month"
	BillWayMissionCount                   BillWay = "mission_count"
	BillWayMissionHold                    BillWay = "mission_hold"
	BillWayMissionVolume                  BillWay = "mission_volume"
	BillWayActiveRegister                 BillWay = "active_register"
	BillWayActiveShare                    BillWay = "active_share"
	BillWayActiveBind                     BillWay = "active_bind"
	BillWayActiveInviteRecharge           BillWay = "active_invite_recharge"
	BillWayActiveChannelRecharge          BillWay = "active_channel_recharge" // 渠道充值赠送
	BillWaySpecialChannelRecharge         BillWay = "special_channel_recharge"
	BillWayActiveRecharge                 BillWay = "active_recharge"
	BillWayTransferManual                 BillWay = "transfer_manual"
	BillWayTransferWithdraw               BillWay = "transfer_withdraw"
	BillWayFirstInviteRecharge            BillWay = "first_invite_recharge"
	BillWayExtraServiceTimePlanHour       BillWay = "extra_service_time_plan_hour"
	BillWayExtraServiceTimePlanDay        BillWay = "extra_service_time_plan_day"
	BillWayExtraServiceTimePlanWeek       BillWay = "extra_service_time_plan_week"
	BillWayExtraServiceTimePlanMonth      BillWay = "extra_service_time_plan_month"
	BillWayExtraServiceHold               BillWay = "extra_service_hold"
	BillWayExtraServiceVolume             BillWay = "extra_service_volume"
	BillWayExtraServiceTime               BillWay = "extra_service_time"
	BillWayWithdrawVX                     BillWay = "withdraw_wechat"
	BillWayWithdrawAlipay                 BillWay = "withdraw_alipay"
	BillWayBankCard                       BillWay = "withdraw_bank_card"
	BillWayWithdrawRefund                 BillWay = "withdraw_refund"
	BillWayWithdrawCompany                BillWay = "withdraw_company" // 对公提现
	BillWayCdkExchange                    BillWay = "cdk_exchange"
	BillWayLottoPrize                     BillWay = "lotto_prize"
	BillWayNodeTrouble                    BillWay = "node_trouble"                      // 节点故障扣费
	BillWayIncomeReplacement              BillWay = "income_replacement"                // 收益补发
	BillWayIncomeDeduct                   BillWay = "income_deduct"                     // 收益扣除
	BillWayIllegalInvite                  BillWay = "illegal_invite"                    // 违规邀请
	BillWaySurveySubmit                   BillWay = "survey_submit"                     // 提交问卷调查
	BillWayCompensateMissionRecoverFailed BillWay = "compensate_mission_recover_failed" // 任务恢复（重新开机）失败补偿
	BillWayInvokedModel                   BillWay = "invoked_model"                     // 调用模型
	BillWayNightSubsidy                   BillWay = "night_subsidy"                     // 夜间补贴
)

func (BillWay) Values() []string {
	return []string{
		string(BillWayBankCard),
		string(BillWayWithdrawAlipay),
		string(BillWayWithdrawVX),
		string(BillWayWithdrawRefund),
		string(BillWayWithdrawCompany),
		string(BillWayUnknown),
		string(BillWayRechargeWechat),
		string(BillWayRechargeAlipay),
		string(BillWayRechargeTron),
		string(BillWayMissionTime),
		string(BillWayMissionTimePlanHour),
		string(BillWayMissionTimePlanDay),
		string(BillWayMissionTimePlanWeek),
		string(BillWayMissionTimePlanMonth),
		string(BillWayMissionCount),
		string(BillWayActiveBind),
		string(BillWayMissionHold),
		string(BillWayMissionVolume),
		string(BillWayActiveRegister),
		string(BillWayActiveShare),
		string(BillWayActiveRecharge),
		string(BillWayTransferManual),
		string(BillWayFirstInviteRecharge),
		string(BillWayTransferWithdraw),
		string(BillWaySpecialChannelRecharge),
		string(BillWayExtraServiceTimePlanHour),
		string(BillWayExtraServiceTimePlanDay),
		string(BillWayExtraServiceTimePlanWeek),
		string(BillWayExtraServiceTimePlanMonth),
		string(BillWayExtraServiceHold),
		string(BillWayExtraServiceVolume),
		string(BillWayActiveInviteRecharge),
		string(BillWayActiveChannelRecharge),
		string(BillWayExtraServiceTime),
		string(BillWayCdkExchange),
		string(BillWayLottoPrize),
		string(BillWayNodeTrouble),
		string(BillWayIncomeReplacement),
		string(BillWayIncomeDeduct),
		string(BillWayIllegalInvite),
		string(BillWaySurveySubmit),
		string(BillWayCompensateMissionRecoverFailed),
		string(BillWayInvokedModel),
		string(BillWayNightSubsidy),
	}
}

func (obj BillWay) Ptr() *BillWay {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
