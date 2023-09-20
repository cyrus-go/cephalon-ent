package enums

type BillWay string

const (
	BillWayUnknown             BillWay = "unknown"
	BillWayRechargeWechat      BillWay = "recharge_wechat"
	BillWayRechargeAlipay      BillWay = "recharge_alipay"
	BillWayMissionTime         BillWay = "mission_time"
	BillWayMissionCount        BillWay = "mission_count"
	BillWayMissionHold         BillWay = "mission_hold"
	BillWayMissionVolume       BillWay = "mission_volume"
	BillWayActiveRegister      BillWay = "active_register"
	BillWayActiveShare         BillWay = "active_share"
	BillWayActiveRecharge      BillWay = "active_recharge"
	BillWayTransferManual      BillWay = "transfer_manual"
	BillWayFirstInviteRecharge BillWay = "first_invite_recharge"
)

func (BillWay) Values() []string {
	return []string{
		string(BillWayUnknown),
		string(BillWayRechargeWechat),
		string(BillWayRechargeAlipay),
		string(BillWayMissionTime),
		string(BillWayMissionCount),
		string(BillWayMissionHold),
		string(BillWayMissionVolume),
		string(BillWayActiveRegister),
		string(BillWayActiveShare),
		string(BillWayActiveRecharge),
		string(BillWayTransferManual),
		string(BillWayFirstInviteRecharge),
	}
}

func (obj BillWay) Ptr() *BillWay {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
