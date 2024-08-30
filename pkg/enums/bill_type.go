package enums

type BillType string

const (
	BillTypeUnknown          BillType = "unknown"
	BillTypeMissionConsume   BillType = "mission_consume"
	BillTypeMissionProduce   BillType = "mission_produce"
	BillTypeRecharge         BillType = "recharge"
	BillTypeTransfer         BillType = "transfer"
	BillTypeActive           BillType = "active"
	BillTypeMission          BillType = "mission"
	BillTypeSpecialMission   BillType = "special_mission" // 特殊任务收益 目前存在分润 但分润为0 不展示给小程序用户
	BillTypeGas              BillType = "gas"
	BillTypeExtraService     BillType = "extra_service"
	BillTypeWithdraw         BillType = "withdraw"
	BillTypeCdk              BillType = "cdk"
	BillTypeLotto            BillType = "lotto"
	BillTypeInvokeModelOrder BillType = "invoke_model_order"
	BillTypeNodeTrouble      BillType = "node_trouble"  // 节点故障扣费
	BillTypeIncomeManage     BillType = "income_manage" // 收益管理
	BillTypeIllegal          BillType = "illegal"       // 违规操作
	BillTypeSurvey           BillType = "survey"        // 问卷调查
	BillTypeCompensate       BillType = "compensate"    // 补偿
	BillTypeSubsidy          BillType = "subsidy"       // 补贴
)

func (BillType) Values() []string {
	return []string{
		string(BillTypeWithdraw),
		string(BillTypeUnknown),
		string(BillTypeRecharge),
		string(BillTypeMissionConsume),
		string(BillTypeMissionProduce),
		string(BillTypeTransfer),
		string(BillTypeActive),
		string(BillTypeMission),
		string(BillTypeSpecialMission),
		string(BillTypeGas),
		string(BillTypeExtraService),
		string(BillTypeCdk),
		string(BillTypeLotto),
		string(BillTypeNodeTrouble),
		string(BillTypeIncomeManage),
		string(BillTypeIllegal),
		string(BillTypeSurvey),
		string(BillTypeCompensate),
		string(BillTypeInvokeModelOrder),
		string(BillTypeSubsidy),
	}
}

func (obj BillType) Ptr() *BillType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
