package enums

type BillType string

const (
	BillTypeUnknown        BillType = "unknown"
	BillTypeMissionConsume BillType = "mission_consume"
	BillTypeMissionProduce BillType = "mission_produce"
	BillTypeRecharge       BillType = "recharge"
	BillTypeTransfer       BillType = "transfer"
	BillTypeActive         BillType = "active"
	BillTypeMission        BillType = "mission"
	BillTypeGas            BillType = "gas"
	BillTypeExtraService   BillType = "extra_service"
	BillTypeWithdraw       BillType = "withdraw"
	BillTypeCdk            BillType = "cdk"
	BillTypeLotto          BillType = "lotto"
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
		string(BillTypeGas),
		string(BillTypeExtraService),
		string(BillTypeCdk),
		string(BillTypeLotto),
	}
}

func (obj BillType) Ptr() *BillType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
