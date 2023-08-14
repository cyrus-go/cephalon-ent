package enums

type BillType string

const (
	BillTypeMissionConsume BillType = "mission_consume"
	BillTypeMissionProduce BillType = "mission_produce"
	BillTypeRecharge       BillType = "recharge"
)

func (BillType) Values() []string {
	return []string{
		string(BillTypeRecharge),
		string(BillTypeMissionConsume),
		string(BillTypeMissionProduce),
	}
}
